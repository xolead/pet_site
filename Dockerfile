FROM golang: 1.23.6 AS builder
WORKDIR /app
COPY handlers.go main.go ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o server
FROM alpine:latest
COPY --from=builder /app/server ./
COPY static/ ./static

EXPOSE 80

CMD ["./server", "--port", "80"]
