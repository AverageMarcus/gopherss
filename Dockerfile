FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.14-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
RUN apk update && apk add --no-cache --update gcc musl-dev git && adduser -D -g '' gopher && apk add -U --no-cache ca-certificates
RUN go get github.com/mattn/go-sqlite3
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o gopherss .

COPY ./views /app/

ENTRYPOINT ["/app/gopherss"]
