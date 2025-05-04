FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /api_gateway cmd/api_gateway/main.go

FROM scratch

COPY --from=builder /api_gateway /api_gateway
COPY --from=builder app/configs/config.yaml /configs/config.yaml

CMD ["/api_gateway"]