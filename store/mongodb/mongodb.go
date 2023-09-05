// TODO:
// - TLS Auth
// - Client options

package mongodb

import (
	"context"
	"errors"
	"log"

	"github.com/tinkershack/muxfarm/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mo "go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	uri    string
	dbName string
	*mongo.Client
}

func Cursor(cur interface{}) (*mongo.Cursor, error) {
	cursor, ok := cur.(*mongo.Cursor)
	if !ok {
		str := "mongodb cursor interface conversion"
		log.Printf("fail: %s", str)
		return nil, errors.New(str)
	}
	return cursor, nil
}

type EmptyOption struct{}

// func (eo *emptyOption) ApplyOption()

type FindOption struct {
	Option *mo.FindOptions
}

// func (fo *findOption) ApplyOption()

type InsertOption struct {
	Option *mo.InsertManyOptions
}

// func (fo *insertOption) ApplyOption()

func (mdb *mongoDB) Option(optionType string) store.DocumentOption {
	switch optionType {
	case "find":
		fo := new(FindOption)
		fo.Option = mo.Find()
		return fo
	case "insert":
		io := new(InsertOption)
		io.Option = mo.InsertMany()
		return io
	default:
		return &EmptyOption{}
	}
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

	serverAPI := mo.ServerAPI(mo.ServerAPIVersion1)
	opts := mo.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("fail: connect MongoDB\n%s", err)
		return client, err
	}

	var result bson.M
	if err := client.Database(m.dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Printf("fail: connect MongoDB\n%s", err)
		return nil, err
	}

	return client, err
}

func (m *mongoDB) Insert(ctx context.Context, collection string, docs []interface{}, option store.DocumentOption) (interface{}, error) {
	col := m.Client.Database(m.dbName).Collection(collection)
	opts := mo.InsertMany().SetOrdered(false)
	ret, err := col.InsertMany(ctx, docs, opts)
	if err != nil {
		log.Printf("fail: insert document\n%s", err)
		return ret, err
	}

	return ret, err
}

func (m *mongoDB) Find(ctx context.Context, collection string, filter interface{}, option store.DocumentOption) (interface{}, error) {
	col := m.Client.Database(m.dbName).Collection(collection)
	fo, ok := option.(*FindOption)
	if !ok {
		log.Fatalf("fail: interface type conversion")
	}
	cursor, err := col.Find(ctx, filter, fo.Option)
	if err != nil {
		log.Printf("fail: Find documents: %s\n", err)
		return cursor, err
	}

	return cursor, nil
}
