## 包

Go 应用程序都是由若干组文件组成，每组文件被称为一个包，也就是 Go 程序里最小的复用单元，可以被其他 Go 程序所使用，比如标准库的http包：
```
net/http
    cgi/
    cookiejar/
        testdata/
    fcgi/
    httptest/
    httputil/
    pprof/
    testdata/
```

> cgi 和 cookiejar 都是一个包，可以单独被引用，使用，他们组成更大的包 http 包

**注意**：同一个目录下所有的 .go 文件都会命名为同一个包名。

## 命名惯例  

给包命名一般都是使用包所在的目录的名字，目录一般都以简洁的小写字母命名。

标准库的实现规范就是这样，这是一个良好实践，我们只需要遵从就可以大大便于协同开发。

## 包的导入

一般使用 `import` 声明块，可以写全路径，明确告知 Go 编译器可以去哪里找到对应的包，也可以写相对路径，那么 Go 编译器就使用 Go 环境变量设置的路径寻找，通常开发者的包就会在 GOPATH 环境变量指定的路径上寻找。

```go
import (
    "fmt"
    "strings"
)
```

## 命名导入

这个特性的好处在于处理同名包的导入问题，对同名的包进行别名命名导入。

```go
package main

import (
    "fmt"
    myfmt "mylib/fmt"
)

func main() {
    fmt.Println("Standard Library")
    myfmt.Println("mylib/fmt")
}
```
