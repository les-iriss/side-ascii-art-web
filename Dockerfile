From golang:1.21.5

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . . 

RUN  go build -o /docker-test

EXPOSE 8080

CMD ["/docker-test"]
################################################

