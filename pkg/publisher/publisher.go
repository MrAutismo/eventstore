package publisher

import (
	"context"

	"github.com/MrAutismo/eventstore/pkg/bootstrap"
	pb "github.com/MrAutismo/eventstore/pkg/pb/eventstorepb/v1"
)

type Publisher struct {
	client pb.EventStoreClient
}

func NewPublisher() *Publisher {
	p := Publisher{
		client: bootstrap.NewEventStoreClient(),
	}

	return &p
}

func (p *Publisher) Publish(ctx context.Context, event *pb.Event) error {
	// t := time.NewTicker(2 * time.Second)
	// go func() {
	// 	for range t.C {
	// event := &pb.Event{
	// 	Id:        uuid.NewString(),
	// 	Name:      "UserRegistered",
	// 	Domain:    "users",
	// 	Data:      []byte("user data in protobuf format"),
	// 	Timestamp: time.Now().Unix(),
	// }

	_, err := p.client.SaveEvent(ctx, &pb.SaveEventRequest{Event: event})
	return err
	// logrus.Infof("res: %+v", res)
	// }
	// }()
}
