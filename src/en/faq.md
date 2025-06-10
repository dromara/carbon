# FAQ

1、What is the difference between `v2.5.x` and `v2.6.x`?
> `v2.5.x` and below use `value` passing, `v2.6.x` and above use `pointer` passing. It is strongly recommended to use `v2.6.x` and above.

2、Timezone error when deploying on `Windows` system

> If the `window` system does not have the `golang` environment installed, the `GOROOT/lib/time/zoneinfo.zip: no such file or directory` exception will be reported during deployment. The reason is that the `window` system does not have a built-in time zone file. You only need to manually download and specify the `zoneinfo.zip` path, such as `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3、Timezone error when deploying on `Docker` container

> If the `docker` container does not have the `golang` environment installed, the `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory` exception will be reported during deployment. You only need to copy `zoneinfo.zip` to the container, that is, add it to the `Dockerfile`

```go
COPY ./zoneinfo.zip /usr/local/go /lib/time/zoneinfo.zip
```