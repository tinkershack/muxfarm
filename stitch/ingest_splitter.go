package stitch

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/tinkershack/muxfarm/config"
	"github.com/tinkershack/muxfarm/dlm"
	"github.com/tinkershack/muxfarm/fixtures"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/store"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO:
// Prevent dangling db entries by cleaning up
// stack cleanup functions before return

func IngestSplitter(ctx context.Context, args []string) {
	config, err := config.New()
	if err != nil {
		log.Fatalf("fail: acquire config\n%s", err)
	}
	log.Printf("config: %+v", config)

	ds, err := mdb.MongoDB(config.MongoDB.URI, config.MongoDB.DBName)
	if err != nil {
		log.Fatalf("fail: acquire MongoDB object\n%s", err)
	}
	defer ds.Client.Disconnect(context.Background())

	dlmc, redisc := dlm.New(config.DLMRedis.URI, config.DLMRedis.DBNumber)
	defer redisc.Close()

	filter := bson.D{{"status.state", "STATE_UNSPECIFIED"}}
	sort := bson.D{{"status.updateTimestamp", 1}}
	opt := ds.Option("find")
	fo, ok := opt.(*mdb.FindOption)
	if !ok {
		str := "mongodb FindOption interface conversion"
		log.Fatalf("fail: %s", str)
	}
	fo.Option.SetSort(sort)
	fo.Option.SetLimit(10)

	for {
		waiter := time.After(15 * time.Second)

		cursor, err := ds.Find(context.Background(), fixtures.IngestCollection, filter, fo)
		if err != nil {
			log.Fatalf("fail: find document: %s", err)
		}
		var docsb []bson.M
		cur, err := mdb.Cursor(cursor)
		if err != nil {
			log.Println("fail: interface cursor conversion: ", err)
			continue
		}
		if err := cur.All(ctx, &docsb); err != nil {
			log.Println("fail: find cursor document: ", err)
			continue
		}

		var wg sync.WaitGroup
		for _, docb := range docsb {
			docb := docb

			wg.Add(1)
			go func() {
				defer wg.Done()

				doc := new(fixtures.IngestDoc)
				err := store.BSONProto(&docb, doc)
				if err != nil {
					log.Printf("fail: BSONProto \n%s", err)
					return
				}
				log.Println("mediain", doc.Mediain.GetInput())

				lock, err := dlmc.Obtain(ctx, doc.Muxfarmid.GetMid(), 1*time.Second, nil)
				if err != nil {
					log.Println("fail: acquire lock:", err)
					return
				}
				defer lock.Release(ctx)
				ticker := time.NewTicker(500 * time.Millisecond)
				defer ticker.Stop()
				go func() {
					for range ticker.C {
						lock.Refresh(ctx, 1*time.Second, nil)
						log.Println("ok: lock refresh key: ", lock.Key())
					}
				}()

				ads := atomDocs(ctx, doc)
				filter := bson.D{{"muxfarmid.mid", doc.Muxfarmid.GetMid()},
					{"status.state", "STATE_UNSPECIFIED"}}
				updateTime, _ := time.Now().UTC().MarshalJSON()
				stateFilter := bson.D{{"$set",
					bson.D{{"status.state", "INGEST_OK"},
						{"status.updateTimestamp", strings.Trim(string(updateTime), "\"")}}}}

				err = ds.FindOneAndUpdate(ctx, fixtures.IngestCollection, filter, stateFilter)
				if err != nil {
					// either no document was found or some thing else went wrong, so bail out
					return
				}

				// Insert AtomDocs
				docsi := make([]interface{}, len(ads))
				for i, v := range ads {
					docb, err := store.ProtoBSON(v)
					if err != nil {
						log.Printf("fail: dangling ingestDoc: %s\n", err)
						return
					}
					docsi[i] = docb
				}
				insOp := ds.Option("")
				ret, err := ds.Insert(context.TODO(), fixtures.AtomCollection, docsi, insOp)
				if err != nil {
					log.Println("fail: insert atomDocs: ", err)
					return
				}
				log.Println("ok: insert atomDocs: ", ret)

				if ttl, err := lock.TTL(ctx); err != nil {
					log.Println("fail: lock ttl fail:", err)
					return
				} else if ttl > 0 {
					log.Println("ok: lock still valid key: ", lock.Key())
					ticker.Stop()
					lock.Release(ctx)
				} else if ttl == 0 {
					log.Println("fail: lock ttl expired early:", err)
					return
				}
				log.Println("AtomDocument count:", len(ads))
			}()
		}
		wg.Wait()
		log.Println("IngestDocument count:", len(docsb))

		<-waiter

	}
}

func atomDocs(ctx context.Context, ingestDoc *fixtures.IngestDoc) []*fixtures.AtomDoc {
	min := ingestDoc.Mediain.GetInput()
	ads := make([]*fixtures.AtomDoc, len(min))
	var wg sync.WaitGroup

	for i, media := range min {
		m := media
		i := i
		wg.Add(1)
		go func(i int, m *plumber.Media) {
			defer wg.Done()
			ad := new(fixtures.AtomDoc)
			mid := new(plumber.MuxfarmID)
			mid.ID()
			ad.Atomid = mid
			ad.Ingestid = ingestDoc.GetMuxfarmid()
			ad.Media = m
			ad.Operation = ingestDoc.Mediain.GetOperation()
			ad.Callerid = ingestDoc.Mediain.GetCallerid()
			ad.Status = plumber.NewStatus(plumber.State_STATE_UNSPECIFIED)
			ad.Timestamp = ad.Status.GetUpdateTimestamp()
			ads[i] = ad
			log.Println("atom doc", i, ad)
		}(i, m)
	}
	wg.Wait()

	return ads
}
