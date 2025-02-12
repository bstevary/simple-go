# Build stage
FROM golang:1.22-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.18
WORKDIR /app 
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY databases/migrations ./databases/migrations

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
