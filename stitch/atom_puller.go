package stitch

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-getter"
	"github.com/tinkershack/muxfarm/dlm"
	"github.com/tinkershack/muxfarm/fixtures"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/store"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO:
// stack errors
// call fail cleanup function before return

func stateAtom(ctx context.Context, atomid string, fromState string, toState string, ds store.Document) error {

	filter := bson.D{{"atomid.mid", atomid},
		{"status.state", fromState}}
	updateTime, _ := time.Now().UTC().MarshalJSON()
	stateFilter := bson.D{{"$set",
		bson.D{{"status.state", toState},
			{"status.updateTimestamp", strings.Trim(string(updateTime), "\"")}}}}

	err := ds.FindOneAndUpdate(ctx, fixtures.AtomCollection, filter, stateFilter)
	if err != nil {
		// either no document was found or some thing else went wrong, so bail out
		return err
	}
	return nil
}

// findAtoms returns Atom Docs of the specified state value
func findAtoms(ctx context.Context, state string, ds store.Document) ([]primitive.M, error) {
	filter := bson.D{{"status.state", state}}
	sort := bson.D{{"status.updateTimestamp", 1}}
	opt := ds.Option("find")
	fo, ok := opt.(*mdb.FindOption)
	if !ok {
		str := "mongodb FindOption interface conversion"
		log.Fatalf("fail: %s", str)
	}
	fo.Option.SetSort(sort)
	fo.Option.SetLimit(fixtures.LimitResultAtomRacer)

	cursor, err := ds.Find(context.Background(), fixtures.AtomCollection, filter, fo)
	if err != nil {
		log.Fatalf("fail: find document: %s", err)
	}
	var docsb []bson.M
	cur, err := mdb.Cursor(cursor)
	if err != nil {
		log.Println("fail: interface cursor conversion: ", err)
		return docsb, err
	}
	if err := cur.All(ctx, &docsb); err != nil {
		log.Println("fail: find cursor document: ", err)
		return docsb, err
	}
	return docsb, nil
}

func atomPull(ctx context.Context, ds store.Document, dlmc *dlm.DLMClient, prCh chan pullResult) {
	for {
		waiter := time.After(15 * time.Second)

		docsb, err := findAtoms(ctx, plumber.State_STATE_UNSPECIFIED.String(), ds)
		if err != nil {
			log.Println("fail: find Atom Docs: ", err)
			return
		}

		getterMode := getter.ClientModeAny
		var wg sync.WaitGroup
		for _, docb := range docsb {
			docb := docb
			wg.Add(1)
			go func() {
				defer wg.Done()

				doc := new(fixtures.AtomDoc)
				err := store.BSONProto(&docb, doc)
				if err != nil {
					log.Printf("fail: BSONProto \n%s", err)
					return
				}
				log.Println("atomid: media", doc.Atomid.GetMid(), doc.GetMedia())

				lock, err := dlmc.Obtain(ctx, doc.Atomid.GetMid(), 5*time.Second, nil)
				if err != nil {
					log.Println("fail: acquire lock:", err)
					return
				}
				defer lock.Release(ctx)
				ticker := time.NewTicker(1 * time.Second)
				defer ticker.Stop()
				go func() {
					for range ticker.C {
						lock.Refresh(ctx, 5*time.Second, nil)
						log.Println("ok: lock refresh key: ", lock.Key())
					}
				}()

				atomURI := doc.Atomid.GetMid() // this var is returned
				atomDir := filepath.Join(fixtures.AtomDirPath, doc.Atomid.GetMid())
				err = os.MkdirAll(atomDir, 0755)
				if err != nil {
					log.Println("fail: create atom dir: ", err)
					return
				}

				sourceURI := doc.Media.FormGetterURI()

				wd, err := os.MkdirTemp("/tmp/", "muxfarm-tmp")
				if err != nil {
					log.Println("fail: create work dir: ", err)
					return
				}
				defer os.RemoveAll(wd)
				getterClient := &getter.Client{
					Ctx:             context.Background(),
					Src:             sourceURI,
					Dst:             wd,
					Pwd:             wd,
					Mode:            getterMode,
					DisableSymlinks: true,
				}

				// Defer call to set atom state
				psFrom := plumber.State_STATE_UNSPECIFIED.String()
				psTo := plumber.State_PULL_FAIL.String()
				psOK := plumber.State_PULL_OK.String()
				var retErr error
				retMedia := new(plumber.Media)
				defer func() {
					stateAtom(ctx, doc.Atomid.GetMid(),
						psFrom,
						psTo,
						ds,
					)
					res := pullResult{
						err: retErr,
						Atom: &fixtures.Atom{
							Atomid: doc.Atomid,
							Media:  retMedia,
						},
					}
					prCh <- res
				}()

				if err := getterClient.Get(); err != nil {
					log.Println("fail: get file: ", err)
					retErr = err
					return
				}
				log.Println("ok: pulled file to ", getterClient.Dst)
				dent, err := os.ReadDir(wd)
				if err != nil {
					log.Println("fail: read work dir: ", err)
					retErr = err
					return
				}
				for _, entry := range dent {
					info, _ := entry.Info()
					if info.Mode().IsRegular() {
						log.Println("ok: file ready for move ", info.Name())
						src := filepath.Join(wd, info.Name())
						dst := filepath.Join(atomDir, info.Name())
						atomURI = filepath.Join(atomURI, info.Name())
						err := os.Rename(src, dst)
						if err != nil {
							log.Println("fail: move file to atom dir: ", err)
							retErr = err
							return
						}
						log.Println("ok: file ready for ops ", dst)
					}
				}

				if ttl, err := lock.TTL(ctx); err != nil {
					log.Println("fail: lock ttl fail:", err)
					retErr = err
					return
				} else if ttl > 0 {
					log.Println("ok: lock still valid key: ", lock.Key())
					ticker.Stop()
					lock.Release(ctx)
					// set variables to be returned by defer fn
					psTo = psOK
					retMedia.Storagetype = *plumber.StorageType_STORAGE_LFS.Enum()
					retMedia.Uri = atomURI
				} else if ttl == 0 {
					log.Println("fail: lock ttl expired early: ", lock.Key(), err)
					retErr = err
					return
				}
			}()
		}

		wg.Wait()
		log.Println("Processed Atom count:", len(docsb))

		<-waiter

	}

}
