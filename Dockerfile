FROM golang:1.22.4 as builder
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ .

RUN go build -o main .
FROM alpine:latest
COPY --from=builder /app/main .
EXPOSE 6666
CMD ["./main"]
