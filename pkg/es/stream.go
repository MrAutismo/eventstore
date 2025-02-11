package es

import (
	"context"

	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *EventStoreServer) StreamInitial(req *pb.StreamEventsRequest, stream pb.EventStore_StreamEventsServer) error {
	ctx := context.Background()
	logrus.Infof("stream inital request: %+v", req)
	var domain interface{} = req.Domain
	if domain == "" {
		domain = bson.M{
			"$regex": ".*",
		}
	}
	var name interface{} = req.Name
	if name == "" {
		name = bson.M{
			"$regex": ".*",
		}
	}

	cursor, err := s.db.Find(ctx, bson.M{"domain": domain, "name": name})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		event, err := s.DecodeCursorEvent(cursor)
		if err != nil {
			return err
		}

		err = s.StreamEvent(stream, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EventStoreServer) StreamEvents(req *pb.StreamEventsRequest, stream pb.EventStore_StreamEventsServer) error {
	ctx := context.Background()

	err := s.StreamInitial(req, stream)
	if err != nil {
		return err
	}
	// Start watching changes in MongoDB
	pipeline := mongo.Pipeline{}
	changeStream, err := s.db.Watch(ctx, pipeline)
	if err != nil {
		return err
	}
	defer changeStream.Close(ctx)

	// Stream new events as they are inserted
	for changeStream.Next(ctx) {
		event, err := s.DecodeChangeEvent(changeStream)
		if err != nil {
			return err
		}

		err = s.StreamEvent(stream, event)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *EventStoreServer) StreamEvent(
	stream pb.EventStore_StreamEventsServer,
	event *pb.Event,
) error {
	logrus.Warnf("sending event: %+v", event)
	// Send event to the client
	return stream.Send(&pb.StreamEventsResponse{Event: event})
}
