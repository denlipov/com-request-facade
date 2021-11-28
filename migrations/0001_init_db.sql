-- +goose Up
create table if not exists requests_events (
        id BIGSERIAL PRIMARY KEY,
        type INT,
        status INT,
        payload JSONB
);

-- +goose Down
drop table requests_events;
