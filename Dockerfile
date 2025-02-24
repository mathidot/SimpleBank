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

# Run stage
FROM alpine:3.19
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .
COPY app.env .

# 暴露端口
EXPOSE 8080 9090

# 设置容器启动命令
CMD [ "/app/main" ]