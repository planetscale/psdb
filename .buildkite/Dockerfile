FROM golang:1.21-bookworm

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y ruby && gem install license_finder && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /src
COPY go.mod go.sum /src
RUN go mod download
