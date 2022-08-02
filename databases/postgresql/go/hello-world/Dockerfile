FROM golang:alpine

WORKDIR /app

COPY example.go .
COPY go.mod .
COPY go.sum .

RUN go build -o example

USER nobody

ENTRYPOINT ["/app/example"]
