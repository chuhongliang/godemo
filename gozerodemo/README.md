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
```


# 启动服务
```bash
// 整理依赖文件
go mod tidy

// 启动 go 程序
$ go run demo.go
```

# 生成api代码
```bash
// 生成api 项目
goctl api new apidemo

// 生成api 代码
goctl api go --api .\user.api --dir .\

// 生成mysql代码
goctl model mysql ddl --src user.sql --dir ../user/
```

# 生成rpc代码
```bash
// 生成rpc 项目
goctl rpc new rpcdemo

// 生成 proto 文件
goctl rpc --o demo.proto


// 生成rpc 代码
goctl rpc proto -o demo.proto --go_out=. --zrpc_out=.

# 单个 rpc 服务生成示例指令
goctl rpc protoc rpcdemo.proto --go_out=. --go-grpc_out=. --zrpc_out=. --client=true

# 多个 rpc 服务生成示例指令
goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --client=true -m
```


# 初始化模版文件
```bash
// 初始化.tpl 模版文件
goctl template init handler.tpl
goctl template init logic.tpl
goctl template init svc.tpl​

// 修改模版文件
vim ~/.goctl/1.7.0/api/handler.tpl
vim ~/.goctl/1.7.0/api/logic.tpl
vim ~/.goctl/1.7.0/api/svc.tpl
```

# swagger 文档
## 编译安装
```bash
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
```

## 配置环境
将$GOPATH/bin中的goctl-swagger添加到环境变量

## 生成swagger.json 文件
```bash
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
```

### 指定Host，basePath，schemes api-host-and-base-path
```bash
goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.1 -basepath /api -schemes https,wss" -api user.api -dir .
```