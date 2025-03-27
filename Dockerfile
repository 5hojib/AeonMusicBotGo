# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

# Initialize module (you can customize the module name)
RUN go mod init aeonbot && \
    go mod tidy && \
    go build -o aeonbot .

# Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/aeonbot .
COPY --from=builder /app/*.dat .

ENV PORT=8080
EXPOSE $PORT

CMD ["./aeonbot"]