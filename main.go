package main

import (
	"log"
	"net"
	"os"

	"github.com/MrAutismo/eventstore/pkg/es"
	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:      true,
			FullTimestamp:    true,
			QuoteEmptyFields: true,
		},
	)
}

func main() {
	server, err := es.NewEventStoreServer()
	if err != nil {
		logrus.Fatalf("Failed to initialize event store: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEventStoreServer(grpcServer, server)

	logrus.Infof("EventStore server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
