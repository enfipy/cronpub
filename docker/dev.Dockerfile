FROM golang:alpine

ENV GO111MODULE on
ARG PROJECT=godev

WORKDIR /go/src/${PROJECT}/

COPY go.mod go.sum ./
RUN apk add git gcc musl-dev && \
    go mod download && \
    go get -u github.com/enfipy/gouto

COPY settings.yaml /

CMD ["gouto", "-dir=src"]
