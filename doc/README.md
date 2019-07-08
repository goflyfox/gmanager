# 跨平台交叉静态编译
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

SET GOROOT=D:\develop\go
SET GOPATH=%GOPATH%;D:\workspace\go_workspace\config-server
go build main.go
main.exe