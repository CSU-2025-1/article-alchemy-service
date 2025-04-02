FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /content_service cmd/content_service/main.go

FROM scratch

COPY --from=builder /content_service /content_service
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/.env /.env

CMD ["/content_service"]