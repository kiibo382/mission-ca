FROM golang:1.15.7-alpine as builder

RUN apk update && \
    apk add --no-cache git mysql-client

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 1
ENV GO111MODULE on
ENV APPDIR $GOPATH/src/mission-ca

WORKDIR ${APPDIR}
COPY go.mod go.sum ${APPDIR}/
RUN go mod download

COPY cmd ${APPDIR}/cmd
COPY pkg ${APPDIR}/pkg
RUN go build -o app ${APPDIR}/cmd/app/main.go

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /app ./
ENTRYPOINT ["./app"]