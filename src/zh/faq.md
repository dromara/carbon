---
head:
  - - meta
    - name: description
      content: 常见问题|一个轻量级、语义化、对开发者友好的 golang 时间处理库,
---

# 常见问题

1、v2.5.x 和 v2.6.x 版本有什么区别？
> `v2.5.x` 及以下版本是`值`传递， `v2.6.x` 及以上版本是`指针`传递，强烈建议使用 `v2.6.x` 及以上版本。

2、window 系统部署时时区报错

> `window` 系统如果没有安装 `golang` 环境，部署时会报 `GOROOT/lib/time/zoneinfo.zip: no such file or directory` 异常，原因是由于
> `window` 系统没有内置时区文件，只需要手动下载并指定 `zoneinfo.zip` 路径即可，如 `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3、docker 容器部署时时区报错

> `docker` 容器如果没有安装 `golang` 环境，部署时会报 `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory`
> 异常，只需要把 `zoneinfo.zip` 复制到容器中即可，即在 `Dockerfile` 中加入

```go
COPY ./zoneinfo.zip /usr/local/go /lib/time/zoneinfo.zip
```