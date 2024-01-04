-- +goose Up
create table feeds (
    id UUID PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url text unique not null,
    user_id UUID NOT NULL references  users(id) on delete cascade
);

-- +goose Down

drop table feeds;
