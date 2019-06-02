## 安装
```shell
$ go get -u github.com/kardianos/govendor
```

## 初始化

在项目根目录下执行进行 vendor 初始化：

```shell
$ govendor init
```

## 从本地（`$GOPATH`）添加依赖
```shell
$ govendor add +e
```

## 拉取依赖
```shell
$ govendor sync
```

---

在实际使用中，windows 平台会报错：Error: CreateFile src: The system cannot find the file specified.

在 `govendor` 的 [issue](https://github.com/kardianos/govendor/issues?utf8=%E2%9C%93&q=The+system+cannot+find+the+file+specified.+) 上也没有很好的解决方法。