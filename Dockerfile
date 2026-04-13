FROM golang:1.26.1-alpine3.23 AS builder

WORKDIR /app
RUN apk update && apk upgrade --no-cache

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache wget && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/api /app/api
COPY migrations ./migrations
USER appuser

EXPOSE 8085

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://127.0.0.1:8080/health || exit 1

CMD ["/app/api"]