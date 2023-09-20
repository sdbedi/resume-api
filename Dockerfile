FROM golang:alpine

WORKDIR /resume-api
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/resume-api/bin/api"]
EXPOSE 8080