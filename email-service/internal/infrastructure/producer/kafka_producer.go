package producer

import (
	"context"
	"email-service/internal/application/service"
	"email-service/internal/pkg"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"

	"time"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func (p *KafkaProducer) Publish(ctx context.Context, payload interface{}) error {
	message, err := p.encodeMessage(payload)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, message)
}

func NewKafkaProducer(brokers []string, topic string) service.Publisher {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: "email-service-client",
	}

	c := kafka.WriterConfig{
		Brokers:          brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}

	return &KafkaProducer{kafka.NewWriter(c)}
}

func (p *KafkaProducer) encodeMessage(payload interface{}) (kafka.Message, error) {
	m, err := json.Marshal(payload)
	if err != nil {
		return kafka.Message{}, err
	}
	//
	key := pkg.NewUUID().String()

	return kafka.Message{
		Key:   []byte(key),
		Value: m,
	}, nil
}