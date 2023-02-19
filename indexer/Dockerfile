# Build stage
FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go build -o indexer .

# Final stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/indexer /app
COPY config.sample.toml /app/config.toml
RUN apk --no-cache add ca-certificates
EXPOSE 8080
CMD ["./indexer", "--config", "/app/config.toml"]
