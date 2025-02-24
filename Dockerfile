# Build stage
FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /app

# 设置代理环境变量
ENV http_proxy=http://172.23.192.1:7890
ENV https_proxy=http://172.23.192.1:7890

# 复制代码到容器中
COPY . .

# 构建 Go 应用程序
RUN go build -mod=readonly -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19
WORKDIR /app
# 从构建阶段复制可执行文件
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x /app/start.sh
COPY wait-for.sh .
RUN chmod +x /app/wait-for.sh
COPY db/migration ./migration

# 暴露端口
EXPOSE 8080 9090
# 设置容器启动命令
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]