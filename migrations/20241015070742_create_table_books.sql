-- +migrate Up
CREATE TABLE "books" (
    "id" bigserial PRIMARY KEY,
    "isbn" text NOT NULL,
    "title" text NOT NULL,
    "author_id" bigserial NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);
ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS "books";