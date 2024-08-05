# 类 Unix
在 Linux 或 macOS 上面，需要运行下面命令（或者，可以把以下命令写到 .bashrc 或 .bash_profile 文件中）：

```bash
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量，以下三选一

# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct

```

确认一下：
```bash
$ go env | grep GOPROXY
GOPROXY="https://goproxy.cn"
```

# Windows
在 Windows 上，需要运行下面命令：
```bash
# 启用 Go Modules 功能
$env:GO111MODULE="on"

# 配置 GOPROXY 环境变量，以下三选一

# 1. 七牛 CDN
$env:GOPROXY="https://goproxy.cn,direct"

# 2. 阿里云
$env:GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"


powershell 官方可用
# 3. 官方
$env:GOPROXY="https://goproxy.io,direct"

# 国内
$env:GOPROXY="https://goproxy.cn,direct"
```
$env:PATH="$PATH:$(go env GOPATH)/bin"


测试一下#

```bash
time go get golang.org/x/tour
```


go install -v golang.org/x/tools/gopls@latest
go install -v github.com/josharian/impl@v1.1.0
go install -v github.com/fatih/gomodifytags@v1.16.0
go install -v github.com/cweill/gotests/gotests@v1.6.0
go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v honnef.co/go/tools/cmd/staticcheck@latest

go install -v install github.com/ramya-rao-a/go-outline@latest
go install -v golang.org/x/lint/golint@latest
go install -v github.com/acroca/go-symbols@latest
go install -v golang.org/x/tools/cmd/guru@latest
go install -v golang.org/x/tools/cmd/gorename@latest

go get -u github.com/mdempsky/gocode
go get -u github.com/sqs/goreturns
go get -u golang.org/x/tools/cmd/gorename


# GO教程
https://learnku.com/docs/the-way-to-go


# Go打包编译
https://developer.aliyun.com/article/1530650


# go build
Mac 下编译Linux, Windows平台的64位二进制文件
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

linux 下编译Mac, Windows平台的64位二进制文件
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

windows 下编译Mac, Linux平台的64位二进制文件
```bash
set CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
set CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

```pm2
pm2 start ./build/server.exe -- dev

```


### Golang 开发笔记
https://www.bookstack.cn/read/golang_development_notes/zh-preface.md

GoAdmin是一个基于 golang 面向生产的数据可视化管理平台搭建框架，可以让你使用简短的代码在极短时间内搭建起一个管理后台
http://doc.go-admin.cn/zh/install/

H2O-Danube3的推出，不仅丰富了开源小型语言模型的生态系统，也为各种应用场景提供了强大的支持。从聊天机器人到特定任务的微调，再到移动设备上的离线应用，H2O-Danube3都展现出了其广泛的适用性和高效性。

模型下载地址：https://top.aibase.com/tool/h2o-danube3

论文地址：https://arxiv.org/pdf/2407.09276


golang 知识合集
https://segmentfault.com/a/1190000038922260


aider ai自动编程
https://aider.chat/docs/usage/modes.html