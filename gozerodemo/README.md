## 配置
```bash
// GO111MODULE 开启
go env -w GO111MODULE=on

// 配置 Proxy​
go env -w GOPROXY=https://goproxy.cn,direct

// 查看配置
$ go env GO111MODULE
on
$ go env GOPROXY
https://goproxy.cn,direct
```

## goctl 配置
```bash
// 安装 goctl
go install github.com/tal-tech/go-zero/tools/goctl@latest

// 查看 goctl 版本
goctl version

// 生成api 项目
goctl api new apidemo

// 生成rpc 项目
goctl rpc new rpcdemo
```

# 启动服务
```bash
// 整理依赖文件
go mod tidy

// 启动 go 程序
$ go run demo.go
```

# 生成代码
```bash
// 生成api 代码
goctl api go --api .\apidemo.api --dir .\


```