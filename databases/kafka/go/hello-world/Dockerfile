FROM golang

WORKDIR /app

COPY go.sum go.mod example.go ./

RUN go build -o kafka-go-hello-world

USER nobody

ENTRYPOINT ["/app/kafka-go-hello-world"]
