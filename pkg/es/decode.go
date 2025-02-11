package es

import (
	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *EventStoreServer) DecodeCursorEvent(stream *mongo.Cursor) (*pb.Event, error) {
	target := &CursorStream{}

	if err := stream.Decode(&target); err != nil {
		return nil, err
	}

	logrus.Infof("decoded event: %+v", target)

	// Deserialize event from BSON
	var event pb.Event
	if err := bson.Unmarshal(target.Data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *EventStoreServer) DecodeChangeEvent(stream *mongo.ChangeStream) (*pb.Event, error) {
	target := &ChangeStream{}

	if err := stream.Decode(&target); err != nil {
		return nil, err
	}

	logrus.Infof("decoded event: %+v", target)

	// Deserialize event from BSON
	var event pb.Event
	if err := bson.Unmarshal(target.FullDocument.Data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
