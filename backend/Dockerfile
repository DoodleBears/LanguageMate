FROM golang:1.22-alpine as builder
WORKDIR /tmp
RUN wget -O gf "https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH)" && \
    chmod +x gf && \
    ./gf install -y && \
    rm ./gf

FROM golang:1.22-alpine as dev
WORKDIR /var/www
COPY --from=builder /go/bin/gf /go/bin/gf
COPY --from=arigaio/atlas:latest-alpine /atlas /usr/local/bin/atlas
COPY manifest/database/migrations /migrations
CMD atlas migrate apply --env docker && \
    gf run main.go

ENV DEBIAN_FRONTEND noninteractive
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai