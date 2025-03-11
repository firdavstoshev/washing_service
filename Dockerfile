FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o washing_service cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/washing_service .
COPY --from=builder /app/config.yml .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./washing_service"]