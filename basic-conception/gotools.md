## 编译
```shell
$ go build hello.go
```

> 在不包含文件名时，直接在目录内执行`go build`，会。

> 还可以指定目录下所有的包，例如：`go build github.com/example/...`。

## 清理
```shell
$ go clean hello.go
```

> `go clean` 与 `go build` 同理。

## 运行
```shell
$ go run hello.go
```

> 该命令会先构建 `hello.go` 里面的程序，然后执行构建后的程序。

## 捕获错误
```shell
$ go vet
```

> 该命令可以开发者检查已有 Go 代码常见的错误，但是不能避免严重的逻辑错误。

## 格式化
```shell
$ go fmt
```

> 该命令能够自动将开发者代码布局成 Go 源代码类似的风格，解决了代码风格统一问题，非常推荐！

> `go fmt` 后面可以跟文件名或者包名。

## 文档
```shell
$ go doc
```

> 使用该命令可以打印文档，即在终端上快速浏览命令或者包的帮助文档，当然还可以借助命令：
```shell
$ godoc -http=:6060
```
> 启动一个 web 服务器，通过 web 形式访问6060端口来浏览本地系统所有 Go 源代码的文档信息。

> 为了让 `godoc` 也能生成文档里包含开发者的代码文档，开发者需要遵从以下规则即可，在标识符前把文档作为注释加入到代码中，这规则对包，函数，类型和全局变量都适用：
```go
// Retry 函数尝试重试
// 成功则返回 true ，失败为 false 并且能够得到 error 信息
func Retry() (bool, error) {
    //...
}
```
