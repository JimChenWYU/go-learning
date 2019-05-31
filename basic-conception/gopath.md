+ GOPATH 环境变量对于 `go install`， `go get` 等 go 工具命令都会用到
+ GOPATH 是作为编译后二进制文件存放位置以及 `import` 其他包时的搜索路径
    - bin
        存放可执行文件
    - src
        存放go代码的源文件
    - pkg
        存放编译好的库文件, 主要是*.a文件
