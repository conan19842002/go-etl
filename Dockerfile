FROM golang:1.21 AS builder

ADD . $GOPATH/app

COPY . .

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl postgresql-client

WORKDIR $GOPATH/app/

COPY --from=builder /app/go-etl .

RUN mkdir -p /app/data/raw /app/data/processed /app/logs

# Download all the dependencies
RUN go get -d -v ./...

RUN go build -ldflags "-s -w" -o etl-pipeline main.go

EXPOSE 6000

CMD ["/app/etl-pipeline"]
