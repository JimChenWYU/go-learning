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