FROM golang:1.12.17-alpine

RUN apk update && \
    apk add git

RUN mkdir /go/src/app
WORKDIR /go/src/app
# COPY ./go.mod .
# COPY ./go.sum .

# RUN go mod download

EXPOSE 8080
