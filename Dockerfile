FROM golang:1.15.5-alpine3.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/maakun12/ramen-search-bot
COPY . .
RUN go build ./cmd/ramen-search-bot/main.go

FROM alpine
COPY --from=builder /go/src/github.com/maakun12/ramen-search-bot /app

CMD /app/main $PORT
