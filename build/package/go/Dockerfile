ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git openssh-client
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /go/src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/go-clean-arch ./cmd/service

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/go-clean-arch /

EXPOSE 8080

ENTRYPOINT ["./go-clean-arch"]
