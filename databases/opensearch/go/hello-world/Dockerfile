FROM golang:alpine

WORKDIR /app

COPY version.go .
COPY go.mod .
COPY go.sum .

RUN go build -o version

USER nobody

ENTRYPOINT ["/app/version"]
