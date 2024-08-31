FROM golang:1.22 AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o main ./main.go

FROM ubuntu:22.04

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
