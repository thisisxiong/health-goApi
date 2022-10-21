#启动编译环境
FROM golang:1.18-alpine AS builder

#配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

#安装mysql

#拷贝源代码
COPY . /go/src/api
WORKDIR /go/src/api

#编译
RUN go build -o app .


FROM mysql:5.7
WORKDIR /docker-entrypoint-initdb.d
#支持utf-8编码机
ENV LANG=C.UTF-8
ENV MYSQL_ROOT_PASSWORD=root
COPY --from=builder /go/src/api/init.sql ./
COPY --from=builder /go/src/api/app ./
RUN mkdir config && chmod 777 app
COPY --from=builder /go/src/api/config/conf.yaml ./config

ENV ADDR=:8080
#暴露端口
EXPOSE 8080

#CMD ["./app"]