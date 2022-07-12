FROM golang:1.17 as builder

LABEL maintainer="keepchen2016@gmail.com"

WORKDIR /build

COPY . /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o drone-feishu

FROM alpine

COPY --from=builder /build/drone-feishu /bin

CMD ["/bin/drone-feishu"]