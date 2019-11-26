FROM golang:1.12.4-stretch AS BUILD

RUN mkdir /dsutils
WORKDIR /dsutils

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD main.go .
ADD dsutils/ ./

RUN go test -v

