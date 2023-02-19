# Build stage
FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go build -o indexer .

# Final stage
FROM ubuntu:22.04
WORKDIR /app
COPY --from=build /app/indexer /app
COPY config.sample.toml /app/config.toml
RUN apt-get update && apt-get install -y ca-certificates
EXPOSE 8080
CMD ["./indexer", "--config", "/app/config.toml"]
