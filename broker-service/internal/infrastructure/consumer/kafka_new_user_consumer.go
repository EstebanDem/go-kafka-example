package consumer

import (
	"broker-service/internal/application/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic string) service.Consumer {
	c := kafka.ReaderConfig{
		Brokers:         brokers,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
		GroupID:         "broker-service",
		StartOffset:     kafka.LastOffset,
	}

	return &KafkaConsumer{kafka.NewReader(c)}
}

func (k *KafkaConsumer) Read(ctx context.Context, chMsg chan service.NewUserMessage, chErr chan error) {
	defer k.reader.Close()

	for {
		m, err := k.reader.ReadMessage(ctx)
		if err != nil {
			chErr <- errors.New("error while reading a message")
			continue
		}

		var message service.NewUserMessage
		err = json.Unmarshal(m.Value, &message)
		if err != nil {
			chErr <- err
		}

		chMsg <- message
	}

}
