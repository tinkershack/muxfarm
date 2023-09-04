package store

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TODO:
// Fix DocumentOption interface methods

type Document interface {
	Option(optionType string) (DocumentOption, error)
	Insert(ctx context.Context, collection string, docs []interface{}, option DocumentOption) (interface{}, error)
	Find(ctx context.Context, collection string, filter interface{}, option DocumentOption) (interface{}, error)
}

type DocumentOption interface {
	// ApplyOption()
}

func ProtoBSON(m protoreflect.ProtoMessage) (*bson.M, error) {
	pj := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	mj, err := pj.Marshal(m)
	if err != nil {
		log.Printf("fail: protojson marshal: %s\n", err)
		return nil, err
	}
	// log.Printf("Ok: protojson message: %+v\n", string(mj))

	var mb bson.M
	if err := json.Unmarshal(mj, &mb); err != nil {
		log.Printf("fail: bson unmarshal: %s\n", err)
		return nil, err
	}

	return &mb, nil
}

func BSONProto(mb *bson.M, pm protoreflect.ProtoMessage) error {
	pj := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	mj, err := json.Marshal(mb)
	if err != nil {
		log.Printf("fail: json marshal: %s\n", err)
		return err
	}
	// log.Printf("Ok: json message: %+v\n", mj)

	if err := pj.Unmarshal(mj, pm); err != nil {
		log.Printf("fail: protojson unmarshal: %s\n", err)
		return err
	}

	return nil
}
