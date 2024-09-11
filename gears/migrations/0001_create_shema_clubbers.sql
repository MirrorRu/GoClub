-- +goose NO Transaction

-- +goose Up
create schema members;


-- +goose Down
drop schema members;
