# 迷你网址导航NavSite (mini navigate site)
- 公司的各种资源太多, 记不住? 这个软件将帮你大忙, 公司所有人员只需要打开这一个网站就可以清楚知道公司内部所有资源
- 一款简单易用的公司内部官网
- 使用golang编写, 只运行一个文件
- 没有第三方依赖, 数据保存为文件, 自己可任何备份
- 前端：vue3 + element
- 后端：golang + gin
- 存储：文件存储，直接存储在服务器上，未使用任何数据库

#### 介绍
使用与公司内部或个人的简易版网址导航工具

#### 示例站点：https://nav.jianean.com/
[demo](https://nav.jianean.com/)
###### 操作账号
- 其它帐号请参见配置文件
- 用户名：add
- 密码：123456
- 权限：只能添加，不能删除和编辑

#### 快速使用
- 选择需要的版本号运行文件, 直接运行文件即可, windows双击, linux./nav-site-server
- 第一次运行时会自动创建默认配置目录和数据目录, 请根据需要修改
- conf目录, 为配置文件, 程序会自动生成conf/config.yaml, 请根据具体需要变更, 暂时不能变更配置位置, 下个版本支持
- conf目录, 如果网站需要不同的favicon, 请将文件放在此目录, 程序会自动加载, 更多的自定义配置敬请期待...
- data目录为数据目录, 1.是JSON数据, 是存储网站导航和分组信息, 2.是网站上传的图片信息, 暂时不能变更配置位置, 下个版本支持
- 当前版本不支持自动备份数据，请自行备份，备份时请同时备份conf和data目录

#### https nginx 配置示例：https://www.jianean.com/blog/74.html
[https nginx 配置示例](https://www.jianean.com/blog/74.html)

#### 安装教程
- 原教程, 新的使用请参见上面的快速使用
- 以下均指Centos7.x服务器

###### 方式一
1. windows构建后端：
   1. 更改默认配置文件：nav-site-server/config/config.yaml 配置如端口号，上传目录，上传图片的访问域名（可以web域名保持一致，斜杠/结尾）
   2. 安装golang >= 1.16
   3. 配置go mod, 安装依赖 go mod tidy
   4. windows编译
   ```
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
   ```
   5. linux系统将程序放置于/opt/nav-site/下, 将nav-site-server.service 中的执行文件路径改成自己的路径，并移动到/etc/systemd/system/nav-site-server.service示例：
```
[Unit]
Description=Nav Site Api Service
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/nav-site/
ExecStart=/opt/nav-site/nav-site-server
Restart=always

[Install]
WantedBy=multi-user.target
```
   8. 执行：systemctl daemon-reload 
   9.  执行：systemctl start nav-site-server.service 启动接口服务
   10. 执行：systemctl status nav-site-server.service 查看接口服务启动状态，或者执行netstat -ntlp | grep :xxxx端口 查看端口是否被监听
   11. 后端完毕


2. 构建前端：（vue3+element）
   1. npm install
   2. npm run build
   3. 上传./dist文件夹中的全部文件到自己的服务器指定的目录，示例：/data/www/nav-site/nav-site-web

3. 配置nginx，
   1. 示例配置：
```
server {
    listen       80;
    server_name  nav.jianean.com;

    #charset koi8-r;
    access_log  /var/log/nginx/nav.jianean.com.access.log  main;
    error_log   /var/log/nginx/nav.jianean.com.error.log;
    error_page   500 502 503 504  /50x.html;
	
    gzip            on;
    gzip_types      text/plain application/xml text/css application/javascript;
    gzip_min_length 1000;

    location / {
        #expires $expires;
        proxy_redirect                      off;
        proxy_set_header Host               $host;
        proxy_set_header X-Real-IP          $remote_addr;
        proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto  $scheme;
        proxy_read_timeout          1m;
        proxy_connect_timeout       1m;
        proxy_pass http://127.0.0.1:xxxx; # gcapi中的config/config.yaml的port端口号
    }
}

```
   3. 执行 nginx -t
   4. 执行 nginx -s reload
   5. 访问域名，查看是否可以正常访问了，如果不能正常访问请自行排查原因，
   6. 作者邮箱: hanlin2531@163.com

###### 方式二
使用已经打包好的文件
-   前端：nav-site-web/dist/ 下的全部文件
-   后端：nav-site-server 下的 ./nav-site-server ./config/config.yaml 一个可执行文件和 config 目录及 config.yaml
-   按照方式一的后面几部配置

#### 开发教程
- 敬请期待

#### 感谢
- 此版本原始设计为https://gitee.com/hyqc/gcwguide, 原作者不再维护, 为了更好的维护, 脱离原来的分支发展