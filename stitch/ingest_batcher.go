package stitch

import (
	"context"
	"log"
	"time"

	"github.com/tinkershack/muxfarm/config"
	"github.com/tinkershack/muxfarm/fixtures"
	"github.com/tinkershack/muxfarm/store"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func IngestBatcher(ctx context.Context, args []string) {
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
			log.Printf("fail: find document: %s", err)
		}
		var docsb []bson.M
		cur, err := mdb.Cursor(cursor)
		if err != nil {
			log.Println("fail: interface cursor conversion: ", err)
		}
		if err := cur.All(ctx, &docsb); err != nil {
			log.Println("fail: find cursor document: ", err)
		}

		var docs []*fixtures.IngestDoc
		for _, docb := range docsb {
			doc := new(fixtures.IngestDoc)
			err := store.BSONProto(&docb, doc)
			if err != nil {
				log.Printf("fail: BSONProto \n%s", err)
			}
			log.Println(doc.Mediain.GetInput())
			docs = append(docs, doc)
		}
		log.Println("IgestDocument count:", len(docs))

		select {
		case <-waiter:

		}
	}
}
