FROM golang:1.16.3-alpine3.13 AS builder
ENV GO111MODULE=on
RUN apk add --no-cache git
WORKDIR /go/src/github.com/MParvin/redis-functionality-exporter
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redis-functionality-exporter .

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/MParvin/redis-functionality-exporter/redis-functionality-exporter /usr/local/bin/redis-functionality-exporter

ENTRYPOINT ["/usr/local/bin/redis-functionality-exporter"]