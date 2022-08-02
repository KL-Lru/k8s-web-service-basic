FROM golang:1.18.4-alpine as builder

WORKDIR /go/src
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN apk update && apk add --no-cache git=2.36.2-r0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd/batch_worker 

FROM alpine:3.16.1

ENV ROOT=/go/app
WORKDIR ${ROOT}

RUN apk update && apk add --no-cache ca-certificates=20211220-r0

COPY --from=builder /go/src/batch_worker ${ROOT}

CMD ["/go/app/batch_worker"]
