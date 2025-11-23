FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo_app main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo_app .

COPY --from=builder /app/web ./web

EXPOSE 7540

CMD ["./todo_app"]