FROM golang:1.13.5-stretch as builder

WORKDIR /app/github.com/zucchinidev/go-ddd-skeleton
COPY . ./
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/policy-api policy/cmd/policy-api/*.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app/github.com/zucchinidev/go-ddd-skeleton
COPY --from=builder /go/bin/policy-api /usr/bin
CMD policy-api

