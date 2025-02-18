-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id UUID NOT NULL PRIMARY KEY,
    email TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
