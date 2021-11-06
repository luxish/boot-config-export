# boot-config-export

Utility CLI to convert [Spring Boot Yaml configuration](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config) into external configuration (as environment variables). The variables are transformed based on the rules described [HERE](https://docs.spring.io/spring-boot/docs/2.5.6/reference/htmlsingle/#features.external-config.typesafe-configuration-properties.relaxed-binding.environment-variables)

## Usage

If you have Go installed you can run:
```
go get github.com/luxish/boot-config-export
go run github.com/luxish/boot-config-export -h
```

## Options

The options for the CLI can be checked by running  `go run github.com/luxish/boot-config-export -h`

| Option      | Default | Description |
|-------------|---------|-------------|
| -f \<path>  | empty   | If specified, the program will read the Yaml file and turn it to an `.env` file.|
| -d \<dir>   | empty   | If specified, the program will read all yaml files and turn them to `.env` files.|
| -o \<output>| "out"   | The output folder for the environment files.|

The configuration is interpreted in this particular order. If the "file" configuration is specified, the "directory" configuration is ignored.
