    -- +goose Up
CREATE TABLE transactions(PRIMARY KEY(id),
    date TEXT,
    transaction TEXT,
);

-- +goose Down
DROP TABLE transactions;