<h1 align="center">我的GO语言学习路线</h1>

> GO语言初学者学习路线规划

## 前言

这是本人初学Go语言的学习路线，总结放上GitHub作为记录，分享

## 目录

- 安装
    + [Linux](install/linux.md)
    + [Windows](install/windows.md)
    + MacOS
- 基本概念
    + [包](basic-conception/package.md)
    + [GOPATH](basic-conception/gopath.md)
    + [Go 工具链](basic-conception/gotools.md)
    + 包管理工具
        - [dep](basic-conception/dep.md)
        - [govendor](basic-conception/govendor.md)
        - [modules](basic-conception/gomodules.md)
- 基本语法
    + [变量定义](grammar/define.go)
    + [init 函数](grammar/init.go)
    + [main 函数](grammar/main.md)
- 数组，切片和映射
    + [数组](array,slice,map/array.go)
    + [切片](array,slice,map/slice.go)
    + [映射](array,slice,map/map.go)
- 类型系统
    + [自定义类型（struct）](typesystem/struct.go)
    + [方法](typesystem/method.go)
    + [类型的本质](typesystem/type.md)
    + [接口](typesystem/interface.go)
    + [嵌入类型](typesystem/typeembedding.go)
    + [标识符](typesystem/identifier.md)
- 并发编程
    + [并发和并行的区别](concurrent/concurrence,parallelism.md)
    + [goroutine](concurrent/goroutine.go)
    + [什么是竞争状态？](concurrent/competition.md)
    + 解决竞争状态的方法
        - [原子函数](concurrent/atomfunction.go)
        - [互斥锁](concurrent/mutex.go)
    + 通道
        - [无缓冲通道](channel/unbuffered_channel.go)
        - [缓冲通道](channel/buffered_channel.go)
    + 并发模式
        - Runner
        - Pool
        - Work
- 标准库
    + 日志
    + JSON编码/解码
    + 输入/输出
- 测试工具
    + 单元测试
        - 基础测试
        - 表组测试
        - 模仿调用
        - 测试服务端点
    + 基准测试
