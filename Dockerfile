FROM golang:1.14-alpine AS builder

ADD . /Users/suryapandian/persona/Go/recipes

WORKDIR /Users/suryapandian/persona/Go/recipes

RUN go build -mod=vendor -o cli .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /Users/suryapandian/persona/Go/recipes . 

CMD ["pwd"]
CMD ["ls"]

ENTRYPOINT [ "./cli", "-search=Chicken", "-file=testData.json", "-time=7AM-12AM"]