[下载 go tools](http://docs.studygolang.com/dl/), 放到`/usr/local`中，并且创建Go的目录`/usr/local/go`，例如：
```shell
$ tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

> 下载源码包时可以根据需要的版本选择下载的文件，比如你想要在 64 位 x86 的Linux 平台上安装 1.2.1 版本的 Go 语言环境，你可以选择 `go1.2.1.linux-amd64.tar.gz`。

（记得运行以上命令时记得赋予 sudo 权限）

添加 `/usr/local/go/bin` 目录到 `PATH` 环境变量中，你可以在 `/etc/profile` 文件 或者 `$HOME/.profile` 中添加如下一行：
```
export PATH=$PATH:/usr/local/go/bin
```

**注意**：变更 `profile` 文件之后有可能不会立刻生效，为了能够马上生效，你应该执行以下命令使其生效：
```shell
$ source $HOME/.profile
```
    