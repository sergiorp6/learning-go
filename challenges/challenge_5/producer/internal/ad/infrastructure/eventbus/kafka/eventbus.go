package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"log"
	"os"
)

type EventBus struct {
	producer sarama.AsyncProducer
	topic    string
}

func NewEventBus(broker string, topic string) *EventBus {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	go func() {
		for msg := range producer.Successes() {
			log.Printf("Event published: %v\n", *msg)
		}
	}()

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to publish event", err)
		}
	}()

	return &EventBus{producer, topic}
}

func (e EventBus) Publish(event domain.Event) error {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return err
	}

	e.producer.Input() <- &sarama.ProducerMessage{
		Topic: e.topic,
		Key:   sarama.StringEncoder(event.AggregateId()),
		Value: sarama.StringEncoder(eventJson),
	}

	return nil
}
