FROM golang:latest
MAINTAINER Razil "JiangtaoShi"

ENV GOPROXY https://goproxy.cn

WORKDIR /src
COPY . .
#go构建可执行文件
RUN go build -o app server.go
#暴露端口
EXPOSE 9091
#最终运行docker的命令
ENTRYPOINT  ["/src/app"]