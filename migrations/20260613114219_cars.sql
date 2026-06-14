-- +goose Up
CREATE TABLE cars (
    id UUID PRIMARY KEY,
    brand TEXT NOT NULL,
    number_plate TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL
);

-- +goose Down
DROP TABLE cars;
