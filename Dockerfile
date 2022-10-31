FROM golang:1.18

WORKDIR /src

COPY . /src

RUN go mod download
