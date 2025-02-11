package streamer

import (
	"context"
	"log"

	"github.com/MrAutismo/eventstore/pkg/bootstrap"
	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
	"github.com/sirupsen/logrus"
)

type Streamer struct {
	client pb.EventStoreClient
}

func NewStreamer() *Streamer {
	s := Streamer{}

	s.client = bootstrap.NewEventStoreClient()

	return &s
}

func (s *Streamer) Stream(ctx context.Context, domain, name string) {
	stream, err := s.client.StreamEvents(
		ctx, &pb.StreamEventsRequest{
			Domain: domain,
			Name:   name,
		},
	)
	if err != nil {
		log.Fatalf("Failed to stream events: %v", err)
	}

	logrus.Info("Listening for events...")
	for {
		resp, err := stream.Recv()
		if err != nil {
			logrus.Fatalf("Stream error: %v", err)
		}
		logrus.Infof("Received event: %+v\n", resp.GetEvent())
	}
}
