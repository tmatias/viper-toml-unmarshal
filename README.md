# viper TOML unmarshal example

```console
$ go run main.go
host=localhost port=5432 user=postgres pass=
output=out.txt
main.Config{Db:main.DatabaseConfig{Host:"localhost", Port:"5432", User:"postgres", Pass:""}, Out:main.OutputConfig{File:"out.txt"}}
```
