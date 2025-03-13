FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init
RUN go mod tidy && go build -o main .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

CMD ["./main"]
