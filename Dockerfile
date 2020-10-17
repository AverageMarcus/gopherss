FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.14-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
RUN apk update && apk add --no-cache git && adduser -D -g '' gopher && apk add -U --no-cache ca-certificates
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o gopherss .

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/gopherss /app/gopherss
COPY ./views /app/views/
USER gopher
ENTRYPOINT ["/app/gopherss"]
