FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o product-service .

EXPOSE 8001

CMD ["./product-service"]
