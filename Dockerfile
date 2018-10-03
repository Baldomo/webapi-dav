FROM golang:1.10.4-alpine as builder
WORKDIR /src/webapi
COPY ./*.go .
RUN GOOS=linux go build -o /out/webapi

FROM alpine:latest
ADD playground/ /webapi-dav/

#COPY tools/startup.sh /webapi-dav/startup.sh

ENV WEBAPI_USER root
ENV WEBAPI_PWD root
ENV MYSQL_ROOT_PASSWORD root

EXPOSE 8080

CMD ["/webapi-dav/startup.sh"]