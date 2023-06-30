FROM golang:latest

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go build main.go

CMD ["./main"]
