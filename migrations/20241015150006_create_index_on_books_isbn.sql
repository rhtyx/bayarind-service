-- +migrate Up
CREATE UNIQUE INDEX "books_isbn_idxkey" ON "books" ("isbn");

-- +migrate Down
DROP INDEX "books_isbn_idxkey";