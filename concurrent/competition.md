### 概念

如果有两个或两个以上的 goroutine 在没有同步的情况下，访问某个共享资源，并试图同时读和写这个资源，就处于互相竞争的状态，这种情况被称作竞争状态。

[例子](./competition.go)
