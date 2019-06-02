## 安装

Go Module 需要 Go 语言版本在 1.11 以上才可以使用，默认没有开启，需要设置环境变量：
```shell
$ export GO111MODULE=on
```

**注意**：在国内使用时，由于从所周知的原因，在 golang.org 上的包是无法下载，那么在使用 Go Module 时同样会遇到该问题，解决方法是可以使用第三方代理或者设置 `GOPROXY`：
```shell
$ export GOPROXY=https://goproxy.io
```

各个平台详细设置可以到 https://goproxy.io

## 初始化

在项目根目录下执行命令：
```shell
$ go mod init

go: creating new go.mod: module xxxx
```

那么你就可以在项目写代码，引入第三方依赖使用，比如在 `main.go`：
```go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

## 构建，执行
```shell
$ go build

$ go run main.go

你好，世界。
```

现在查看项目目录下 `go.mod` 文件
```shell
$ cat go.mod

module xxxx

require rsc.io/quote v1.5.2
```

查看 `$GOPATH/pkg` 目录会发现，多了 `mod` 目录，进去可以看到：
```
cache
golang.org
rsc.io
    quote@v1.5.2
    sampler@v1.3.0
```

那么我们就明白了，当我们执行构建时，会自动识别到我们项目中依赖哪些第三方包，然后下载到 `$GOPATH/pkg/mod` 中并缓存，构建时会自动包含，所以使用 `Go Module` 时 `import` 包时包的路径只需要明确包名即可
