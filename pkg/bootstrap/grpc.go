package bootstrap

import (
	"log"

	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewEventStoreClient() pb.EventStoreClient {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := pb.NewEventStoreClient(conn)

	return client
}
