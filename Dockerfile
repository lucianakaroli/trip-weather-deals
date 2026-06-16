# Etapa de build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN go build -o trip-weather-deals ./cmd/api

# Etapa final
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/trip-weather-deals .

EXPOSE 8080

CMD ["./trip-weather-deals"]