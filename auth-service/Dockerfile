FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o auth-service .

EXPOSE 8000

CMD ["./auth-service"]
