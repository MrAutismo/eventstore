package es

import (
	"context"
	"fmt"

	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *EventStoreServer) SaveEvent(ctx context.Context, req *pb.SaveEventRequest) (*pb.SaveEventResponse, error) {
	logrus.Infof("Received SaveEvent request: %v", req)
	event := req.GetEvent()

	if event == nil {
		return &pb.SaveEventResponse{Success: false}, fmt.Errorf("invalid event")
	}
	eventBson, err := bson.Marshal(event)
	if err != nil {
		return &pb.SaveEventResponse{Success: false}, err
	}
	_, err = s.db.InsertOne(
		ctx,
		bson.M{
			"_id":       event.Id,
			"data":      eventBson,
			"name":      event.Name,
			"domain":    event.Domain,
			"timestamp": event.Timestamp,
		},
	)
	if err != nil {
		return &pb.SaveEventResponse{Success: false}, err
	}
	return &pb.SaveEventResponse{Success: true}, nil
}
