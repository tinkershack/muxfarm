package mops

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/store"
	"google.golang.org/protobuf/encoding/protojson"
)

// Implements mops service defined in Plumber proto
type mops struct {
	plumber.UnimplementedMopsServer
	ds store.Document // Document Store
	// mu sync.Mutex
}

func Mops(db store.Document) *mops {
	m := new(mops)
	m.ds = db
	return m
}

type ingestDoc struct {
	*plumber.MuxfarmID
	*plumber.MediaIn
}

func (m *mops) Ingest(ctx context.Context, min *plumber.MediaIn) (*plumber.MuxfarmID, error) {
	mid := plumber.NewMuxfarmID(uuid.NewString())
	doc := new(ingestDoc)
	doc.MuxfarmID = mid
	doc.MediaIn = min
	pj := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	docj, err := pj.Marshal(doc)

	docs := make([]interface{}, 1)
	docs[0] = doc
	log.Printf("%+v\n%+v", min, docs[0])
	ret, err := m.ds.Insert(context.TODO(), "ingest", docs)
	if err != nil {
		log.Printf("Fail: insert document: %s\n", err)
		return nil, err
	}
	log.Printf("Ok: insert id: %+v", ret)

	return mid, nil
}
