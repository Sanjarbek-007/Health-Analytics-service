package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type ConsumerKafka interface {
	ComsumeMessages(handler func(message []byte)) error
	Close() error
}

type ConsumerKafkaImpl struct {
	reader *kafka.Reader
}

func NewKafkaConsumInit(broker []string, topic string, groupID string) (*ConsumerKafkaImpl, error) {
	return &ConsumerKafkaImpl{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: broker,
			GroupID: groupID,
			Topic:   topic,
		}),
	}, nil
}

func (c *ConsumerKafkaImpl) ComsumeMessages(handler func(message []byte)) error {
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		handler(m.Value)
	}
}

func (c *ConsumerKafkaImpl) Close() error {
	return c.reader.Close()
}
