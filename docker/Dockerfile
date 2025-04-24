# 使用基础镜像
FROM ubuntu:2204-ali-source
RUN apt-get update && apt-get install -y nginx
ENV LANG="en_US.UTF-8"
RUN rm /etc/nginx/sites-enabled/default
RUN rm -rf /usr/share/nginx/html*
COPY web/dist/ /usr/share/nginx/html/
RUN rm -rf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 设置工作目录
RUN mkdir -p  /data/app
WORKDIR /data/app
# 复制配置文件
COPY ./configs configs
# 复制启动脚本
COPY ./server server
# 复制依赖包
ENV GO_ENV=prod
# 公开端口
EXPOSE 8000
EXPOSE 8001

# 使用 ENTRYPOINT 同时启动 nginx 和 server
ENTRYPOINT ["sh", "-c", "nginx -g 'daemon off;' & ./server"]
