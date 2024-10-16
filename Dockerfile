FROM golang:1.22.8 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./main

RUN make cert

FROM debian:jessie-slim

COPY --from=builder /app/main /
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/cert /cert/
COPY --from=builder /app/config.yml /

EXPOSE 8010

CMD ["sleep", "infinity"]