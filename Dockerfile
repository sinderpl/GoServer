
# -------------------    GO SERVER   ---------------------------

FROM golang:1.17-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

COPY ./app/*.go ./

RUN go get -d ./...

RUN go build -o /go-server

EXPOSE 8080

CMD [ "/go-server" ]

# -------------------    GO SERVER   ---------------------------