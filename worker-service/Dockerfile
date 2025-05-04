FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /worker-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
