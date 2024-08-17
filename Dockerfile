FROM golang:1.22.6-alpine
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest