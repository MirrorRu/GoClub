# Postgres db migrations
+goose NO TRANSACTION

-- +goose Up
-- +goose StatementBegin
create schema members;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
create schema members;
-- +goose StatementEnd