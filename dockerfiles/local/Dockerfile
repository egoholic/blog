FROM golang:1.13.0-alpine3.10 as stoa-blog

WORKDIR /go/src/app

COPY . /go/src/app
RUN apk add git
EXPOSE 3000
RUN go get -v ./...
CMD ["go", "run", "/go/src/app/targets/web/main.go"]
