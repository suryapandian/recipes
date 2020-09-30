FROM golang:1.15-alpine AS builder

ADD . /go/src/github.com/hellofreshdevtests/suryapandian-recipe-count-test-2020

WORKDIR /go/src/github.com/hellofreshdevtests/suryapandian-recipe-count-test-2020

RUN go build -mod=vendor -o cli .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/github.com/hellofreshdevtests/suryapandian-recipe-count-test-2020 . 

ENTRYPOINT [ "./cli"]