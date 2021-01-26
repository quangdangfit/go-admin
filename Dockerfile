FROM golang:alpine as builder

MAINTAINER quangdp<quangdangfit@gmail.com>

WORKDIR /go/src/go-admin
COPY . /go/src/go-admin
RUN go build -o ./dist/go-admin

FROM alpine:3.11.3
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

COPY ./config/config.yaml .
COPY --from=builder /go/src/go-admin/dist/go-admin .
EXPOSE 8888
ENTRYPOINT ["./go-admin"]
