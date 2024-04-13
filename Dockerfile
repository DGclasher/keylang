FROM golang:1.22.2-alpine

WORKDIR /app
COPY . /app

RUN go build -o bin/keylang cmd/main.go

CMD ["/app/bin/keylang"]