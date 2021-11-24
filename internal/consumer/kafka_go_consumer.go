package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/denlipov/com-request-facade/internal/db"
	"github.com/denlipov/com-request-facade/internal/model"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type kafkaConsumer struct {
	db     db.Storage
	reader *kafka.Reader
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func NewEventConsumer(
	brokers []string,
	topic string,
	group string,
	db db.Storage) (Consumer, error) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: group,
		Topic:   topic,
	})

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	consumer := &kafkaConsumer{
		db:     db,
		reader: reader,
		cancel: cancel,
		wg:     wg,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := consumer.reader.ReadMessage(ctx)
			if err != nil {
				ctxErr := ctx.Err()
				if ctxErr != nil || errors.Is(err, io.EOF) {
					log.Error().Err(err).Msg("Complete consumer")
					return
				} else {
					continue
				}
			}
			consumer.processMessage(ctx, msg)
		}
	}()

	return consumer, nil
}

func (c *kafkaConsumer) Close() {
	c.cancel()
	c.wg.Wait()
}

func (c *kafkaConsumer) processMessage(ctx context.Context, msg kafka.Message) {

	var event model.RequestEvent

	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Error().Err(err).Msg("Message unmarsal failed")
		return
	}

	log.Debug().Msgf("processing message from kafka, eventID: %d", event.ID)

	err = c.db.DumpMessage(ctx, event)
	if err != nil {
		log.Error().Err(err).Msg("DB saving message failed")
	}
}
