FROM golang:1.23.0-alpine
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest