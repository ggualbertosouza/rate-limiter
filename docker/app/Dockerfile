FROM golang:1.25.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./

RUN go build -o api ./cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/api ./

RUN chmod +x ./api

CMD ["./api"]