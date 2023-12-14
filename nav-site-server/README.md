## 后端调试
- 安装golang >= 1.18
- 配置go mod, 安装依赖 go mod tidy

### 前端调试
```
直接调试
修改前端vue.config.js中配置的地址或端口
修改后端的启动端口
(以上可不修改,如果本地不冲突)
单独运行前端代码和后端代码即可开始调试

将前端打包进go语言调试
将前端打包好的代码, 目录一般在nav-site-web/dist
复制到/nav-site-server/server/static下, 
运行后端代码即可

快速部署前端代码到golang工程里, 进入script目录, 运行脚本, 
如windows: PS E:\2project\golang\goland\001\NavSite\nav-site-server\script> .\install-npm.bat
```

### 后台打包
```
右键运行main.go即可
```

### 后台打包
```
打包环境说明
GOOS: linux,windows,DARWIN,FREEBSD
GOARCH: amd64,386,arm

编译linux
go env -w CGO_ENABLED=0
go env -w GOOS=LINUX
go env -w GOARCH=amd64
go build -ldflags '-s -w'

编译windows
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags '-s -w'

1. Windows主机编译Linux,MAC客户端
Windows主机编译Windows客户端
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags '-s -w' 或 go build -o nav-site-server.exe main.go

Windows主机编译LINUX客户端
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -ldflags '-s -w' 或 go build -o nav-site-server main.go

Windows主机编译MAC客户端
go env -w CGO_ENABLED=0
go env -w GOARCH=darwin
go env -w GOARCH=amd64
go build -ldflags '-s -w' 或 go build -o nav-site-server main.go

2 Linux主机编译Widows,MAC客户端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nav-site-server main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o nav-site-server main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o nav-site-server.exe main.go

3 MAC主机编译Widows,linux客户端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nav-site-server main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o nav-site-server main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o nav-site-server.exe main.go

```

#### 发布至github
```
发布本地版本
goreleaser release --auto-snapshot --rm-dist

git tag -a v0.0.18 -m 'release v0.0.18'
git push origin v0.0.18

发布到github上
goreleaser release --rm-dist
```

### 调试失效
```bash
go get -u github.com/go-delve/delve/cmd/dlv@latest
```
