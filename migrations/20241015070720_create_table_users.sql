-- +migrate Up
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" text UNIQUE NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "users";