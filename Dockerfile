FROM golang:1.21.5

WORKDIR /app

COPY . .

RUN go build -o main . 

CMD ["./main"]
