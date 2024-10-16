-- +migrate Up
CREATE TABLE "authors" (
    "id" bigserial PRIMARY KEY,
    "name" text NOT NULL,
    "birth_date" timestamp NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "authors";