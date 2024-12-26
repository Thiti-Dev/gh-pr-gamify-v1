FROM golang:1.23.1 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux
    # GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

COPY ./pkg ./pkg

RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /app/config.yaml

CMD ["/app/main"]