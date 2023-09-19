FROM golang:alpine

WORKDIR /resume-api
COPY . .

RUN go build main.go

CMD ["/resume-api/main"]
EXPOSE 8080