FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY m3db-go-hello.go .
RUN go build -o m3db-go-hello-world

USER nobody

ENTRYPOINT ["/app/m3db-go-hello-world"]
