package postgres

import (
	"context"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-facade/internal/model"
	"github.com/jmoiron/sqlx"
)

// pgEventRepo ...
type pgEventRepo struct {
	db *sqlx.DB
}

// NewEventRepo ...
func NewEventRepo(db *sqlx.DB) *pgEventRepo {
	return &pgEventRepo{
		db: db,
	}
}

func (r *pgEventRepo) DumpMessage(ctx context.Context, event model.RequestEvent) error {

	req := event.Entity
	payload, err := json.Marshal(*req)
	if err != nil {
		return err
	}

	// RequestEvent
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.db).
		Insert("requests_events").
		Columns(
			"type",
			"status",
			"payload").
		Values(
			event.Type,
			event.Status,
			payload)

	_, err = query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
