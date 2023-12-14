# 迷你网址导航NavSite (mini navigate site)
- 公司的各种资源太多, 记不住? 这个软件将帮你大忙, 公司所有人员只需要打开这一个网站就可以清楚知道公司内部所有资源
- 一款简单易用的公司内部官网
- 使用golang编写, 只运行一个文件
- 没有第三方依赖, 数据保存为文件, 自己可任何备份
- 前端：vue3 + element
- 后端：golang + gin
- 存储：文件存储，直接存储在服务器上，未使用任何数据库

## 介绍
使用与公司内部或个人的简易版网址导航工具

## 示例站点
[demo](http://nav.cifaz.com:8083/)

#### 操作账号
- 其它帐号请参见配置文件
- 用户名：add
- 密码：123456
- 权限：只能添加，不能删除和编辑

## 快速使用
- 安装go语言环境, 请参与网上, 此处不列出, 安装好后, 请检测
```
go version 
版本应该在1.16及以上
```
- 下载最新程序, 地址: [https://github.com/cifaz/nav-site/releases](https://github.com/cifaz/nav-site/releases)
- 第一次运行时会自动创建默认配置目录和数据目录, 请根据需要修改
- windows安装, 下载到你指定的目录解压后, 双击exe文件运行即可, 默认端口号:8083, 请使用http://ip:8083访问, 如需指定配置目录,请使用命令行运行, 具体参见linux参数
- linux安装, 下载到指定目录后, 
```
# 1. 安装golang环境包, 如有请跳过
yum -y install golang wget
# 2. 下载运行包
cd /opt/ && wget https://github.com/cifaz/nav-site/releases/download/v0.0.18/nav-site_0.0.18_Linux_x86_64.tar.gz

# 3. 解压安装包
tar xzvf nav-site_0.0.18_Linux_x86_64.tar.gz --strip-components 1 -C nav-site
mv nav-site_0.0.18_Linux_x86_64 nav-site

# 4. 赋执行权限
chmod u+x /opt/nav-site/nav-site

# 5. 运行程序
普通运行
/opt/nav-site/nav-site

指定配置目录运行 
/opt/nav-site/nav-site conf-dir=/opt/nav-site/

后台运行
nohup /opt/nav-site/nav-site > /opt/nav-site/nav-site.out 2>&1 &

# 6. 登录创建数据即可
默认用户密码: admin 123456
```
- linux服务化 systemctl
```
示例脚本仅适用于centos7/8, 其它系统请参考
# 复制conf目录下的nav-site-server-centos7-8.service至并命名为/etc/systemd/system/nav-site.service
cp /opt/nav-site/conf/nav-site.service /etc/systemd/system/nav-site.service

# 重新载入配置
systemctl daemon-reload 
# 启动程序
systemctl start nav-site
# 查看启动状态
systemctl status nav-site.service 
# 查看端口是否被监听
netstat -ntlp | grep :8083

```
- 程序配置介绍
```
conf目录, 为配置文件, 程序会自动生成conf/config.yaml, 请根据具体需要变更
conf目录, 如果网站需要不同的favicon, 请将文件放在此目录, 程序会自动加载, 更多的自定义配置, 敬请期待...
data目录为数据目录, 1.是JSON数据, 是存储网站导航和分组信息, 2.是网站上传的图片信息
```
- 当前版本不支持自动备份数据，请自行备份，备份时请同时备份conf和data目录

## 联系作者交流
作者邮箱: hanlin2531@163.com

## 感谢
- 此版本原始设计为https://gitee.com/hyqc/gcwguide, 原作者不再维护, 为了更好的维护, 脱离原来的分支发展
