FROM --platform=linux/amd64 golang:1.21.5-alpine3.19

WORKDIR usr/app

COPY . .

RUN go get -d -v ./...

RUN go build -o ./bin/api ./cmd/api

EXPOSE 8000

CMD ["./bin/api"]
