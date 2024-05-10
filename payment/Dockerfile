FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o payment_bin main.go

CMD ["./payment_bin"]