FROM golang:1.23.7-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth_service cmd/auth_service/main.go

FROM scratch

COPY --from=builder /auth_service /auth_service
COPY --from=builder /app/migrations /migrations
COPY --from=builder /app/.env /.env

CMD ["/auth_service"]