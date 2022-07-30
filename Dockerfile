#builder
FROM golang:1.17.7-alpine3.15 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api main.go

#runner 
FROM alpine:3.15

COPY --from=builder /app/api /app/

WORKDIR /app

EXPOSE 8080

CMD ["./api"]