FROM golang:alpine as builder

WORKDIR /go/src/github.com/clementlecorre/mail-to-telegram
ADD . .

RUN apk add --no-cache ca-certificates tzdata
RUN go build

FROM alpine

ENV TZ=Europe/Paris

WORKDIR /go/src/github.com/clementlecorre/mail-to-telegram

COPY --from=builder /usr/share/zoneinfo/ /usr/share/zoneinfo/
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/github.com/clementlecorre/mail-to-telegram/mail-to-telegram .
ENTRYPOINT ["./mail-to-telegram"]
