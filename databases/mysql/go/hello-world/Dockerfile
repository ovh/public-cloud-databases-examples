FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY version.go .
RUN go build -o version

USER nobody

ENTRYPOINT ["/app/version"]
