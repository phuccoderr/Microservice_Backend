FROM golang:1.22.5-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o mailService ./cmd/server/

FROM scratch

COPY ./config /config

COPY ./internal/kafka/templates /internal/kafka/templates

COPY --from=builder /build/mailService /mailService

ENTRYPOINT ["./mailService", "config/local.yaml"]