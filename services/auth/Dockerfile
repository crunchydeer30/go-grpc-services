FROM golang:1.23.2-alpine AS builder

COPY . /grpc-services/auth
WORKDIR /grpc-services/auth

RUN go mod download
RUN go build -o ./bin/server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /grpc-services/auth/bin/server .

CMD ["./server"]