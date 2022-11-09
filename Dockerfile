FROM golang:1.19 AS builder

WORKDIR /go/src/nats
COPY . .
RUN go mod vendor

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -ldflags="-extldflags=-static" -a -installsuffix nocgo -tags=nomsgpack -o /app main.go

FROM debian:buster-slim


COPY --from=builder /app ./


ENTRYPOINT ["./app"]
