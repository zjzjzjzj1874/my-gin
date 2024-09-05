

## openapi

* docs: swagger文档
* 生成客户端的client
```shell
  openapi-generator generate -i swagger.json -g go --additional-properties=packageName=myclient,withGoMod=false -o ../openapi/mygin
#  openapi-generator generate -i swagger.json -g go --additional-properties=packageName=myclient,withGoMod=false,GIT_USER_ID=my-gin,GIT_REPO_ID=openapi -o ../openapi/mygin
```

* [调用demo](../examples/demo.go)
