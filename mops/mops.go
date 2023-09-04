package mops

import (
	"context"
	"log"

	"github.com/tinkershack/muxfarm/fixtures"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/store"
)

// Implements mops service defined in Plumber
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

// TODO:
// Validate fields before ingest

func (m *mops) Ingest(ctx context.Context, min *plumber.MediaIn) (*plumber.MuxfarmID, error) {
	doc := new(fixtures.IngestDoc)
	mid := new(plumber.MuxfarmID)
	mid.ID()
	doc.Muxfarmid = mid
	doc.Mediain = min
	doc.Status = plumber.NewStatus(plumber.State_STATE_UNSPECIFIED)
	doc.Timestamp = doc.Status.GetUpdateTimestamp()

	docb, err := store.ProtoBSON(doc)
	if err != nil {
		log.Printf("Fail: ProtoBSON: %s\n", err)
		return nil, err
	}

	docs := make([]interface{}, 1)
	docs[0] = docb
	log.Printf("\nproto: %+v\nbson: %+v\n", doc, docs[0])
	insOp, _ := m.ds.Option("")
	ret, err := m.ds.Insert(context.TODO(), fixtures.IngestCollection, docs, insOp)
	if err != nil {
		log.Printf("Fail: insert document: %s\n", err)
		return nil, err
	}
	log.Printf("Ok: insert id: %+v\n", ret)

	return mid, nil
}
