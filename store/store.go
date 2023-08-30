package store

import "context"

type Document interface {
	Insert(ctx context.Context, collection string, docs []interface{}) (interface{}, error)
}
