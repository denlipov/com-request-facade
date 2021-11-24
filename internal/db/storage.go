package db

import (
	"context"

	"github.com/denlipov/com-request-facade/internal/db/postgres"
	"github.com/denlipov/com-request-facade/internal/model"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	DumpMessage(ctx context.Context, event model.RequestEvent) error
}

func NewStorage(db *sqlx.DB) Storage {
	return postgres.NewEventRepo(db)
}
