FROM golang:1.24

WORKDIR /app

ARG FILE

COPY ${FILE} main.go

RUN go build -o app main.go

CMD ["./app"]
