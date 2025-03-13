FROM golang:1.18-alpine

ADD . app/
# Set the Current Working Directory inside the container
WORKDIR app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl postgresql-client


RUN mkdir -p /app/data/raw /app/data/processed /app/logs

RUN CGO_ENABLED=0 GOOS=linux go build -o etl-pipeline main.go

EXPOSE 6000

CMD ["/etl-pipeline"]
