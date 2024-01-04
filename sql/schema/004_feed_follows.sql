-- +goose Up
create table feed_follows (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id),
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table feed_follows;