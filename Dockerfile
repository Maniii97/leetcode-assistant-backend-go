FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

# Set environment to production to skip .env file loading
ENV APP_ENV=production

COPY --from=builder /app/main .

CMD ["./main"]
