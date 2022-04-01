# boot-config-export

Utility CLI to convert [Spring Boot YAML configuration](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config) into external configuration (as environment variables). The variables are transformed based on the rules described [here](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config.typesafe-configuration-properties.relaxed-binding.environment-variables).


## Usage

With Go version 1.16+:

```bash
go install github.com/luxish/boot-config-export@latest

go run github.com/luxish/boot-config-export@latest -h
```

File to export:
```yml
# application.yaml
spring:
  profiles:
    active: dev
  main:
    banner-mode: off
  server: 
    port: 9999
```

```
> go run github.com/luxish/boot-config-export -f application.yaml
SPRING_MAIN_BANNERMODE=false 
SPRING_PROFILES_ACTIVE="dev"
SPRING_SERVER_PORT=9999
```

```
> go run github.com/luxish/boot-config-export -f application.yaml -t cm
apiVersion: v1
kind: ConfigMap
metadata:
  name: "ExportedConfig"
  labels: {}
  annotations: {}
data:
  SPRING_MAIN_BANNERMODE: "false"
  SPRING_PROFILES_ACTIVE: "dev"
  SPRING_SERVER_PORT: "9999"
```
