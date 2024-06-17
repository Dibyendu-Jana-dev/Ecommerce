FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o order-service .

EXPOSE 8002

CMD ["./order-service"]
