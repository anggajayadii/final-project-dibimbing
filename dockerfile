# Stage 1: Build
FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Stage 2: Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy executable
COPY --from=builder /app/main .

# Copy .env file
COPY --from=builder /app/.env .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]
