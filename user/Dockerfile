FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o user_bin main.go

CMD ["./user_bin"]