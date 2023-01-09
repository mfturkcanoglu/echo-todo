CREATE TABLE IF NOT EXISTS todo (
    id bigserial PRIMARY KEY,
    text varchar NOT NULL,
    created_at timestamptz DEFAULT (now()),
    deleted boolean
);