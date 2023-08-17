FROM golang:1.20-buster AS builder

WORKDIR /var/app

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go clean -modcache ; go mod tidy

RUN go build -o /pismo_challenge

EXPOSE 80

ENTRYPOINT ["/pismo_challenge"]