FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on \
GOPROXY=https://proxy.golang.org,direct \
GOMODCACHE=/go/pkg/mod

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/aeonbot .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/bin/aeonbot .
COPY --from=builder /app/*.dat .

ENV PORT=8080
EXPOSE $PORT

CMD ["./aeonbot"]