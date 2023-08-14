FROM golang:1.16-buster AS builder

WORKDIR /var/app

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go build -o /pismo_challenge

EXPOSE 8080

ENTRYPOINT ["/pismo_challenge"]