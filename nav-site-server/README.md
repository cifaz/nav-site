## 后端调试
### 前端
```
将前端打包好的代码, 一般在nav-site-web/dist
复制到/nav-site-server/server/static下, 

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

发布到github上
goreleaser release --rm-dist
```

