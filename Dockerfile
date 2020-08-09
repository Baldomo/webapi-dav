FROM golang:alpine as builder

RUN apk add --no-cache git

WORKDIR /src/

COPY build.go .
COPY ./pkg ./pkg
COPY ./cmd ./cmd
COPY go.mod go.sum .
RUN go get ./...

RUN GOOS=linux CGO_ENABLED=0 go run build.go -fast -os linux build
RUN apk del git

FROM scratch
COPY orario.xml /
COPY config.toml /
COPY ./docs /docs
COPY --from=builder /src/build/linux/webapi /

ENV WEBAPI_DB_USER root
ENV WEBAPI_DB_PWD root

EXPOSE 8080:8080

CMD ["/webapi"]
