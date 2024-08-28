FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

RUN chmod a+x ./server

CMD ["./server"]