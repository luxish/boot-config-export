# boot-config-export

Utility CLI to convert [Spring Boot YAML configuration](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config) into external configuration (as environment variables). The variables are transformed based on the rules described [here](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config.typesafe-configuration-properties.relaxed-binding.environment-variables).


## Usage

If Go version 1.16+ is installed installed:

```bash
go install github.com/luxish/boot-config-export@latest

go run github.com/luxish/boot-config-export@latest -h

go run github.com/luxish/boot-config-export@latest -f example/test.yaml
```


| <div style="width:100px">Option</div> | Default | Description |
|---------------|---------|-------------|
| -f \<path>    | empty   | The program will read the YAML file and will export the configuration in the desired format. |
| -t \<type>    | env     | If specified the output will be changed based on the type. Options: **env** (environment variables), **cm** (K8s ConfigMap resource) |
| -o \<output>  | empty   | Output file name. |
