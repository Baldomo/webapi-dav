FROM golang:alpine as builder

RUN apk add --no-cache git sh

RUN mkdir -p /go/webapi
WORKDIR /go/webapi

COPY build.go .
COPY ./internal ./internal
COPY ./cmd ./cmd
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk del git

RUN GOOS=linux CGO_ENABLED=0 go run build.go -fast -os linux build

FROM scratch
RUN mkdir -p /go/bin/webapi
COPY orario.xml /go/bin
COPY config.toml /go/bin
COPY ./docs /go/bin/docs
COPY --from=builder /go/bin/webapi /go/bin/webapi

ENV WEBAPI_DB_USER root
ENV WEBAPI_DB_PWD root
ENV MYSQL_ROOT_PASSWORD root

EXPOSE 8080:8080

CMD ["/go/bin/webapi"]