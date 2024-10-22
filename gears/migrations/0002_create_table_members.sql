
-- +goose Up
-- +goose NO TRANSACTION
CREATE SEQUENCE IF NOT EXISTS order_id_manual_seq INCREMENT 100 START 100 ;

CREATE table IF NOT EXISTS members.members (
    id INT DEFAULT nextval('order_id_manual_seq') PRIMARY KEY,  -- задаём автоматическое назначение идентификаторов
    full_name VARCHAR(100) NOT NULL,
    birth_day date,
    created_at TIMESTAMPTZ DEFAULT NOW()
);


-- +goose Down
DROP table IF EXISTS members.members;

DROP SEQUENCE IF EXISTS order_id_manual_seq;
