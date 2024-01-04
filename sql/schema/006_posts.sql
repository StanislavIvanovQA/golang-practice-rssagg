-- +goose Up
create table posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE ,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table posts;
