## dep 安装

### Linux (Ubuntu)
```shell
sudo apt-get install go-dep
```

### Windows
在 [go.equinox.io](https://go.equinox.io/github.com/golang/dep/cmd/dep) 上下载压缩包解压，通过 `powershell` 或者 `cmd` 运行 dep.exe

### MacOS
```shell
$ brew install dep

$ brew upgrade dep
```

### Other Platform
```shell
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```


## 创建新项目
```shell
$ mkdir -p $GOPATH/src/github.com/me/example

$ cd $GOPATH/src/github.com/me/example
```

## 初始化
```shell
$ dep init

$ ls
Gopkg.toml Gopkg.lock vendor/
```

## 添加第三方依赖
```shell
$ dep ensure -add github.com/bitly/go-simplejson@=0.4.3 #这里 @= 参数指定的是 某个 tag

$ dep ensure #同步
```

## 删除没有使用的依赖
```shell
$ dep prune
```

## 使用时遇到的问题

+ 当项目内没有Go的代码文件时，安装第三方包会报错 `no dirs contained any Go code`。
+ 当项目内没有引用安装的第三方包时，安装时会提示 `"xxxxxx" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.`。
+ 项目下有两个由 `dep` 生成的文件：`Gopkg.toml` 和 `Gopkg.lock`，注意 `Gopkg.toml` 可以手动编辑，而 `Gopkg.lock` 不能随意更改。

## [例子项目](depproject)
