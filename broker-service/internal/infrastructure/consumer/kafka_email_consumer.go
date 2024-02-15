package consumer

import (
	"broker-service/internal/application/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaEmailConsumer struct {
	reader *kafka.Reader
}

func NewKafkaEmailConsumer(brokers []string, topic string) service.ConsumerEmail {
	c := kafka.ReaderConfig{
		Brokers:         brokers,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		GroupID:         "broker-service-email",
		StartOffset:     kafka.LastOffset,
	}

	return &KafkaEmailConsumer{kafka.NewReader(c)}
}

func (k *KafkaEmailConsumer) Read(ctx context.Context, chMsg chan service.EmailMessage, chErr chan error) {
	defer k.reader.Close()

	for {
		m, err := k.reader.ReadMessage(ctx)
		if err != nil {
			chErr <- errors.New("error while reading a message")
			continue
		}

		var message service.EmailMessage
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			chErr <- err
		}

		chMsg <- message
	}

}
