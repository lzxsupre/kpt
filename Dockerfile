FROM golang:alpine

ENV TZ=Asia/Shanghai

RUN mkdir -p /go/src/github.com/mivinci/kpt
COPY . /go/src/github.com/mivinci/kpt
WORKDIR /go/src/github.com/mivinci/kpt

RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go build cmd/main.go
ENTRYPOINT [ "./main" ]

