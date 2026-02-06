FROM golang:1.25.6

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build main.go

CMD ["./main"]