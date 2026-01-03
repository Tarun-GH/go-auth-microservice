
#---------Build Stage---------
FROM golang:alpine AS builder

WORKDIR /app

# copy go.mod files first (cache optimization)
COPY go.mod go.sum ./
RUN go mod download 

# copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server/main.go

#----------Runtime stage----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/app .

# Expose app port
EXPOSE 8080

# Run the binary
CMD ["./app"]