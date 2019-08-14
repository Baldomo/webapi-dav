FROM golang:alpine as builder

RUN apk add --no-cache git sh

RUN mkdir -p /go/webapi
WORKDIR /go/webapi

COPY ./internal ./internal
COPY ./cmd ./cmd
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk del git

RUN mkdir -p /out
RUN GOOS=linux CGO_ENABLED=0 go build -o /go/bin/webapi ./cmd/webapi

FROM scratch
RUN mkdir -p /go/bin/webapi
COPY orario.xml /go/bin
COPY config.toml /go/bin
COPY ./static /go/bin/static
COPY --from=builder /go/bin/webapi /go/bin/webapi

ENV WEBAPI_USER root
ENV WEBAPI_PWD root
ENV MYSQL_ROOT_PASSWORD root

EXPOSE 8080:8080

CMD ["/go/bin/webapi"]