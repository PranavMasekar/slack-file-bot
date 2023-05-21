FROM golang:1.16-alpine As builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o go-slack-file .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/go-slack-file /app/go-slack-file

CMD ["/app/go-slack-file"]