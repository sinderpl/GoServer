# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-server

EXPOSE 8080

CMD [ "/go-server" ]