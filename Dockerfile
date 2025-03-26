FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o aeonmusicbot .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/aeonmusicbot .
COPY --from=builder /app/*.dat .

CMD ["./aeonmusicbot"]