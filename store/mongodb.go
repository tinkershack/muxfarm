// TODO:
// - TLS Auth
// - Client options

package store

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tinkershack/muxfarm/plumber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	uri    string
	dbName string
	*mongo.Client
}

func MongoDB(uri, dbName string) (*mongoDB, error) {
	mdb := new(mongoDB)
	mdb.uri = uri
	mdb.dbName = dbName
	var err error
	mdb.Client, err = mdb.connect(context.TODO())
	if err != nil {
		log.Printf("Fail: acquire MongoDB\n%s", err)
		return nil, err
	}

	return mdb, nil
}

func (m *mongoDB) connect(ctx context.Context) (*mongo.Client, error) {
	uri := m.uri

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("Fail: connect MongoDB\n%s", err)
		return client, err
	}

	var result bson.M
	if err := client.Database(m.dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Printf("Fail: connect MongoDB\n%s", err)
		return nil, err
	}
	log.Printf("Out: Ping: %v", result)

	return client, err
}

func (m *mongoDB) Insert(ctx context.Context, collection string, docs []interface{}) (interface{}, error) {
	col := m.Client.Database(m.dbName).Collection(collection)
	opts := options.InsertMany().SetOrdered(false)
	ret, err := col.InsertMany(ctx, docs, opts)
	if err != nil {
		log.Printf("Return: %+v\n", ret)
		log.Printf("Fail: insert document\n%s", err)
		return ret, err
	}

	type ingestDoc struct {
		*plumber.MuxfarmID
		*plumber.MediaIn
	}
	// doc := new(ingestDoc)
	// min := plumber.NewMediaIn()
	// cid := &plumber.CallerID{Cid: "tc2"}
	// min.Cid = cid
	// doc.MediaIn = min
	// log.Printf("Min: %+v", doc)

	fil := bson.D{{"mediain.cid.cid", "tc2"}}
	cur, e := col.Find(ctx, fil)
	if err != nil {
		log.Printf("Fail: find document\n%s", e)
	}
	var res []ingestDoc
	if e := cur.All(ctx, &res); e != nil {
		log.Printf("Fail: find cursor document\n%s", e)

	}

	for _, result := range res {
		res, _ := json.Marshal(result)
		log.Println(string(res))
	}

	log.Printf("Out.find: %+v\n%v", res, len(res))
	return ret, err
}
