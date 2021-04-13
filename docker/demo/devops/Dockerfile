FROM golang:1.16.3 AS builder
WORKDIR /go/src/gitee.com/jianzhuliu/docker
COPY main.go .
ENV GO111MODULE auto
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/http-server .

FROM scratch
WORKDIR /root/
EXPOSE 8004
COPY --from=builder /tmp/http-server . 
CMD ["./http-server"]
