package es

import (
	"context"

	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventStoreServer struct {
	pb.UnimplementedEventStoreServer
	db *mongo.Collection
}

type ChangeStream struct {
	FullDocument struct {
		ID   string `bson:"_id"`
		Data []byte `bson:"data"`
	} `bson:"fullDocument"`
}

type CursorStream struct {
	Data []byte `bson:"data"`
}

func NewEventStoreServer() (*EventStoreServer, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/replicaSet=rs0"))
	if err != nil {
		return nil, err
	}
	db := client.Database("eventstore").Collection("events")
	return &EventStoreServer{db: db}, nil
}
