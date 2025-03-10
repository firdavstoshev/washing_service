FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o washing_service cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/washing_service .
EXPOSE 8080
CMD ["./washing_service"]