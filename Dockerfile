 # -------- Build Stage --------
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main -ldflags="-s -w" main.go

# -------- Runtime Stage --------
FROM alpine:3.8

ENV WORKDIR=/app
WORKDIR $WORKDIR

COPY --from=builder /app/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

CMD ["./main"]