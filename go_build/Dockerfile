FROM golang:1.14.6-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base
RUN set -x \
    && go get -u github.com/graphql-go/graphql \
    && go get -u github.com/graphql-go/handler \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/postgres 
WORKDIR /go/src