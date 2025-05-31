-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY ,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text NOT NULL,
    url text UNIQUE NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);



-- +goose Down
DROP TABLE feeds;
