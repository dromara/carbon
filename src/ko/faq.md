---
head:
  - - meta
    - name: description
      content: "Carbon 자주 묻는 질문, 버전 차이 설명(v2.5.x 와 v2.6.x), Windows 시스템 타임존 오류 해결 방법, Docker 컨테이너 배포 타임존 설정 등 일반적인 사용 문제 포함"
  - - meta
    - name: keywords
      content: "carbon, go-carbon, 자주 묻는 질문, FAQ, 버전 차이, 타임존 오류, Windows 배포, Docker 배포, 타임존 설정, 문제 해결, 트러블슈팅"
---

# 자주 묻는 질문

1. v2.5.x와 v2.6.x 버전의 차이점은 무엇인가요?
> `v2.5.x` 이하 버전은 `값` 전달이고, `v2.6.x` 이상 버전은 `포인터` 전달입니다. `v2.6.x` 이상 버전 사용을 강력히 권장합니다.

2. Windows 시스템 배포 시 시간대 오류

> `Windows` 시스템에 `golang` 환경이 설치되어 있지 않으면 배포 시 `GOROOT/lib/time/zoneinfo.zip: no such file or directory` 예외가 발생합니다. 이는
> `Windows` 시스템에 내장 시간대 파일이 없기 때문이며, `zoneinfo.zip`을 수동으로 다운로드하고 경로를 지정하면 됩니다. 예: `go/lib/time/zoneinfo.zip`

```go
os.Setenv("ZONEINFO", "./go/lib/time/zoneinfo.zip")
```

3. Docker 컨테이너 배포 시 시간대 오류

> `Docker` 컨테이너에 `golang` 환경이 설치되어 있지 않으면 배포 시 `open /usr/local/go/lib/time/zoneinfo.zip: no such file or directory`
> 예외가 발생합니다. `zoneinfo.zip`을 컨테이너에 복사하면 됩니다. 즉, `Dockerfile`에 다음을 추가하세요

```dockerfile
COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
``` 