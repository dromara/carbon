---
head:
  - - meta
    - name: description
      content: 프로젝트 소개|가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리
---

# 프로젝트 소개

`Carbon` 은 가벼우면서도 의미론적이고 개발자 친화적인 `golang` 시간 처리 라이브러리로, `어떤` 서드파티 라이브러리에도 의존하지 않으며, `100%` 단위 테스트 커버리지를 가지고 있으며, [docker](https://github.com/docker/docker-language-server/blob/main/go.mod#L10 "docker")에 공식적으로 사용되고 있고 [awesome-go](https://github.com/yinggaozhen/awesome-go-cn#日期和时间 "awesome-go-cn") 와 [hello-github](https://hellogithub.com/repository/dromara/carbon "hello-github")에도 수록되어 있습니다.

<img src="/docker.jpg" width="100%" alt="docker"/>

## 프로젝트 특징
- 가벼움 & 제로 의존성: 단위 테스트를 제외하고는 `어떤` 서드파티 라이브러리에도 의존하지 않으며, `100%` 단위 테스트 커버리지로 코드 품질과 안정성을 보장합니다
- 의미론적 API: 의미론적이고 개발자 친화적인 API를 제공하며, 간결하고 우아한 체이닝 호출로 코드의 확장성과 재사용성을 보장합니다
- 강력한 시간 조작: 시간 파싱, 시간 더하기/빼기, 시간 설정, 시간 경계, 시간 판단, 시간 차이, 시간 극값 등을 지원합니다
- 풍부한 가져오기 방식: 시간의 각 부분(년, 월, 일, 시, 분, 초 등)과 다양한 정밀도의 타임스탬프를 가져옵니다
- 다양한 출력 형식: 필요에 따라 다양한 정밀도와 형식의 시간 문자열을 출력하며, ISO, RFC 시리즈 및 사용자 정의 형식을 포함합니다
- 테스트 친화적: 테스트 시간 설정을 지원하며, 현재 시간을 동결하여 단위 테스트에 편리합니다
- 시간대 처리: 시간대 설정 및 가져오기를 지원하며, 서로 다른 시간대 간의 상호 변환을 지원합니다
- 국제화: 30개 이상의 현지화 번역 언어를 지원하며 사용자 정의 번역 리소스를 허용합니다
- 오류 처리: 오류 검사 메커니즘을 제공하여 시간 파싱 등의 오류를 쉽게 처리할 수 있습니다

## 저장소 주소

[github.com/dromara/carbon](https://github.com/dromara/carbon "github.com/dromara/carbon")

[gitee.com/dromara/carbon](https://gitee.com/dromara/carbon "gitee.com/dromara/carbon")

[gitcode.com/dromara/carbon](https://gitcode.com/dromara/carbon "gitcode.com/dromara/carbon")