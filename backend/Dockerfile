# --- Stage 1: Build binary
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Static build
RUN CGO_ENABLED=0 GOOS=linux go build -o haiku-server .

# --- Stage 2: Minimal runtime
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/haiku-server /

USER nonroot:nonroot

ENTRYPOINT ["/haiku-server"]