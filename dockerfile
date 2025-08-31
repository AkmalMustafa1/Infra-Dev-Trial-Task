FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o solana-api ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/solana-api .
ENV PORT=8080
EXPOSE 8080
CMD ["./solana-api"]
