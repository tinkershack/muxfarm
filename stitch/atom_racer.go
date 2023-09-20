package stitch

import (
	"context"
	"log"

	"github.com/tinkershack/muxfarm/config"
	"github.com/tinkershack/muxfarm/dlm"
	"github.com/tinkershack/muxfarm/fixtures"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
)

// TODO:
// - stack cleanup functions before return

type pullResult struct {
	*fixtures.Atom
	err error
}

func AtomRacer(ctx context.Context, args []string) {
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

	// pull result chan
	prCh := make(chan pullResult, fixtures.LimitResultAtomRacer)

	go atomPull(ctx, ds, dlmc, prCh)

	for {
		select {
		// Check for result from atomPull
		case res := <-prCh:
			if res.err != nil {
				log.Println(err)
			} else {
				log.Println("ok AtomPullResult:", res.Atom.Atomid.GetMid(), res.Atom.Media.GetUri())
			}
		}
	}

}
