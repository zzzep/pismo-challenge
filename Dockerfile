FROM golang:1.20-buster AS builder

WORKDIR /var/app

COPY . .

RUN go mod download

RUN go clean -modcache

RUN go mod tidy

RUN go build -o /pismo_challenge

EXPOSE 80

ENTRYPOINT ["/pismo_challenge"]