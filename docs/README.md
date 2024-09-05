

## openapi

* docs: swagger文档
* 生成客户端的client
```shell
  openapi-generator generate -i swagger.json -g go --additional-properties=packageName=myclient,withGoMod=false -o ../openapi/mygin
```

* [调用demo](../examples/demo.go)
