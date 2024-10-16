-- +migrate Up
CREATE TABLE "sessions" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigserial NOT NULL,
    "refresh_token" text NOT NULL,
    "refresh_token_expired_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp
);
ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS "sessions";