package stdout

import (
	"context"

	"github.com/denlipov/com-request-facade/internal/model"
	"github.com/rs/zerolog/log"
)

type stdoutStorage struct {
}

func NewStorage() *stdoutStorage {
	return &stdoutStorage{}
}

func (s *stdoutStorage) DumpMessage(ctx context.Context, event model.RequestEvent) error {
	log.Info().Msgf("Dump message from kafka: %s", event.String())
	return nil
}
