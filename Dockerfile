FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8081

CMD ["./main"]