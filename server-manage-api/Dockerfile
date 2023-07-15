FROM golang:1.20 AS builder

WORKDIR /app
ADD . /app

RUN go build -o server-manager-api cmd/server/main.go

FROM ubuntu:latest AS launcher
COPY --from=builder /app .
CMD ["./server-manager-api"]