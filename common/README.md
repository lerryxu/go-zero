## 1. 配置代理
```
go env -w GOPROXY=https://goproxy.cn,direct
```
## 2. 安装goctl
```
GO111MODULE=on go install github.com/zeromicro/go-zero/tools/goctl@1.5.4
```
## 3. 安装protoc
```
goctl env check --install --verbose --force
```
