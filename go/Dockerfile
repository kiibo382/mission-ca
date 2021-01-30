FROM golang:latest

WORKDIR /go/src

RUN apt-get update \
  && apt-get install git \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

COPY ./src /go/src

EXPOSE 8000

ENTRYPOINT [ "go", "run" , "main.go" ]