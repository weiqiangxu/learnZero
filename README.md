# learnZero

gozero实现单体服务

# 启动服务

```
cd greet &&  go run greet.go -f etc/greet-api.yaml
```

#访问服务

```
$ curl -i -X GET http://localhost:8888/from/you
```