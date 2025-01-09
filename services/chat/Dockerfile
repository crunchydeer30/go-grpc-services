FROM golang:1.23.2-alpine AS builder

COPY . /grpc-services/chat
WORKDIR /grpc-services/chat

RUN go mod download
RUN go build -o ./bin/server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /grpc-services/chat/bin/server .

CMD ["./server"]