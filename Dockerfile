FROM golang:1.11.2-stretch

RUN apt-get update

RUN apt-get install -y \
  default-libmysqlclient-dev \
  mysql-client

WORKDIR /go/src/github.com/mochizukikotaro/go-repository-pattern

RUN go get github.com/oxequa/realize
RUN go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate
RUN go get -u github.com/golang/dep/cmd/dep