FROM golang:1.18.4-alpine as builder

WORKDIR /go/src
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN apk update && apk add --no-cache git=2.36.2-r0 ca-certificates=20211220-r0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd/batch_worker 

FROM scratch

ENV ROOT=/go/app
WORKDIR ${ROOT}

COPY --from=builder /go/src/batch_worker ${ROOT}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/go/app/batch_worker"]
