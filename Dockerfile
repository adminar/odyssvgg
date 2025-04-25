# === 阶段 1：构建后端 Go 应用 ===
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o odyssey cmd/server/main.go

# === 阶段 2：构建前端 Vue3 应用 ===
FROM node:22-alpine AS frontend-builder
WORKDIR /web
COPY web/package*.json ./
RUN npm install
COPY web/ .
RUN npm run build 

# === 阶段 3：生产环境镜像 ===
FROM nginx:alpine

# 安装证书
RUN apk add --no-cache ca-certificates

# 设置工作目录
WORKDIR /usr/bin

# 复制前端构建结果到 nginx 静态目录
COPY --from=frontend-builder /web/dist/ /usr/share/nginx/html/

# 复制自定义 nginx 配置（待补充）
COPY ./nginx.conf /etc/nginx/nginx.conf

# 复制后端二进制
COPY --from=backend-builder /app/odyssey /usr/bin/odyssey

# 复制后端配置
COPY --from=backend-builder /app/configs /usr/bin/configs

# 开放端口
EXPOSE 80
EXPOSE 8000

# 启动后端和 nginx
CMD ["sh", "-c", "/usr/bin/odyssey & nginx -g 'daemon off;'"]