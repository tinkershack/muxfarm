package sense

import (
	"context"
	"errors"
	"log"

	"github.com/tinkershack/muxfarm/fixtures"
	"github.com/tinkershack/muxfarm/store"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func IngestBatcher(ctx context.Context, ds store.Document) error {
	filter := bson.D{{"status.state", "STATE_UNSPECIFIED"}}
	sort := bson.D{{"status.updateTimestamp", 1}}
	opt, err := ds.Option("find")
	if err != nil {
		log.Fatalf("fail: doc store option: %s", err)
	}
	fo, ok := opt.(*mdb.FindOption)
	if !ok {
		str := "mongodb FindOption interface conversion"
		log.Printf("fail: %s", str)
		return errors.New(str)
	}
	fo.Option.SetSort(sort)
	fo.Option.SetLimit(10)
	cursor, err := ds.Find(context.Background(), fixtures.IngestCollection, filter, fo)
	if err != nil {
		log.Printf("fail: find document: %s", err)
	}
	var docsb []bson.M
	cur, err := mdb.Cursor(cursor)
	if err != nil {
		log.Printf("fail: interface cursor conversion: %s", err)
	}
	if err := cur.All(ctx, &docsb); err != nil {
		log.Printf("fail: find cursor document\n%s", err)

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

	return errors.New("fail: IngestBatcher")
}
