---
head:
  - - meta
    - name: description
      content: Carbon FAQ, covering version differences (v2.5.x vs v2.6.x), Windows system timezone error solutions, Docker container deployment timezone configuration and other common usage issues
  - - meta
    - name: keywords
      content: carbon, go-carbon, FAQ, frequently asked questions, version differences, timezone error, Windows deployment, Docker deployment, timezone configuration, troubleshooting, problem solving
---

# FAQ

1. What is the difference between `v2.5.x` and `v2.6.x`?
> `v2.5.x` and below use `value` passing, `v2.6.x` and above use `pointer` passing. It is strongly recommended to use `v2.6.x` and above.

2. Timezone error when deploying on `Windows` system

> If the `Windows` system does not have the `golang` environment installed, the `GOROOT/lib/time/zoneinfo.zip: no such file or directory` exception will be reported during deployment. The reason is that the `Windows` system does not have a built-in time zone file. You only need to manually download and specify the `zoneinfo.zip` path, such as `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3. Timezone error when deploying on `Docker` container

> If the `docker` container does not have the `golang` environment installed, the `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory` exception will be reported during deployment. You only need to copy `zoneinfo.zip` to the container, that is, add it to the `Dockerfile`

```dockerfile
COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
```