-- +goose Up
CREATE TABLE posts (
    id uuid UNIQUE PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    url TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id uuid NOT NULL
);

-- +goose down
DROP TABLE posts;