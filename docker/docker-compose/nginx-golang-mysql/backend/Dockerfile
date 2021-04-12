FROM golang:1.14.15-alpine3.12 AS build
WORKDIR /go/src/github.com/org/repo
COPY . .

ENV GOPROXY https://goproxy.cn,direct
RUN go env
RUN go build -o server .

FROM alpine:3.12
EXPOSE 8000
COPY --from=build /go/src/github.com/org/repo/server /server
CMD ["/server"]
