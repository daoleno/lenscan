# Build stage
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o indexer .

# Final stage
FROM ubuntu
RUN apt update && apt install ca-certificates -y
WORKDIR /app
COPY --from=build /app/indexer .
CMD ["./indexer"]
