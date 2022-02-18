# FROM golang:1.16-alpine AS builder
FROM arm64v8/golang:1.16

RUN apk add --update build-base
ENV GOPROXY=https://goproxy.io
RUN mkdir /tmp/app && \
apk add curl stress-ng
COPY . /tmp/app/
RUN cd /tmp/app && \
go mod vendor

RUN cd /tmp/app && \ 
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags="-w -s"  -o /app ./

RUN ls /tmp/app 

CMD [ "/app"]