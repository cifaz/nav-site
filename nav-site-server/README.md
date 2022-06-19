### 开发
```
开发前, 请先执行web打包, 并复制到/nav-site-server/server/static下, 
windows下可使用快速执行脚本
cd script
./install-npm.bat

```

### 打包
```
打包神器 
go get github.com/mitchellh/gox

打包
GOOS: linux,windows,DARWIN,FREEBSD
GOARCH: amd64,386,arm

linux编译
go env -w CGO_ENABLED=0
go env -w GOOS=LINUX
go env -w GOARCH=amd64
go build -ldflags '-s -w'

windows编译
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -ldflags '-s -w'

goreleaser release --auto-snapshot --rm-dist
goreleaser release --snapshot --rm-dist
goreleaser release --skip-publish --skip-validate  --rm-dist
goreleaser build --single-target --rm-dist

暂时没有实现直接发布到github
暂时只有Windows环境打包脚本

```

#### 其它
```

2.1 Windows主机编译Linux,MAC客户端
Windows主机编译Windows客户端

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o abc-demo-windows.exe main.go
Windows主机编译LINUX客户端

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o abc-demo-linux main.go
Windows主机编译MAC客户端

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o abc-demo-mac main.go
2.2 Linux主机编译Widows,MAC客户端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abc-demo-linux main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o adc-demo-mac main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o abc-demo-windows.exe main.go
2.3 MAC主机编译Widows,linux客户端
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abc-demo-linux main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o abc-demo-mac main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o abc-demo-windows.exe main.go
```
