###############################################################################
#                                   SERVER
###############################################################################

# 第一阶段：编译go项目
FROM golang:1.24.4-alpine as server-builder
ENV WORKDIR                 /app
WORKDIR $WORKDIR

COPY ./server .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# 第二阶段：部署server项目
FROM alpine:3.21.3 as gmanager-server

ENV WORKDIR                 /app
# 复制编译后的静态文件和启动文件
COPY --from=server-builder /app/resource $WORKDIR/resource
COPY --from=server-builder /app/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 暴露端口
EXPOSE 8000

WORKDIR $WORKDIR
CMD ./main


###############################################################################
#                                   NGINX
###############################################################################

# 第一阶段：编译pnpm项目
FROM node:22.16.0-alpine as nginx-builder
# 安装pnpm
RUN npm install -g pnpm
# 设置镜像
RUN pnpm config set registry https://registry.npmmirror.com/
# 复制项目文件
WORKDIR /app
COPY ./web .
# 安装依赖并编译
RUN rm -rf node_modules
RUN pnpm install
RUN pnpm build

# 第二阶段：使用Nginx部署
FROM nginx:1.28.0-alpine as gmanager-nginx
# 复制编译后的静态文件
COPY --from=nginx-builder /app/dist /app/dist
# 暴露端口
EXPOSE 80
# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]