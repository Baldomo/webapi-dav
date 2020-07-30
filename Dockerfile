FROM golang:alpine as builder

RUN apk add --no-cache git

RUN mkdir -p /go/webapi
WORKDIR /go/webapi

COPY build.go .
COPY ./pkg ./pkg
COPY ./cmd ./cmd
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk del git

RUN GOOS=linux CGO_ENABLED=0 go run build.go -fast -os linux build

# FROM scratch
RUN mkdir -p /go/bin
WORKDIR /go/bin
COPY orario.xml .
COPY config.toml .
COPY ./docs ./docs
# COPY --from=builder /go/bin/build/linux/webapi /go/bin/webapi
RUN cp /go/webapi/build/linux/webapi .

# ENV WEBAPI_DB_USER root
# ENV WEBAPI_DB_PWD root
ENV MYSQL_ROOT_PASSWORD root

EXPOSE 8080:8080

CMD ["/go/bin/webapi"]