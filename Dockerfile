FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY templates/ templates/
COPY static/ static/

EXPOSE 8000
CMD ["./main"]