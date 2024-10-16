-- +migrate Up
CREATE UNIQUE INDEX "sessions_refresh_token_idxkey" ON "sessions" ("refresh_token");

-- +migrate Down
DROP INDEX "sessions_refresh_token_idxkey";