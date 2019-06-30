`counters`

```go
package counters

// 未公开的类型
type alertCounter int

// 公开的方法
func New(value int) alertCounter {
    return alertCounter(value)
}
```

`main`

```go
package main

import (
    "fmt"
    "counters"
)

func main() {
	counter := counters.New(10)
    fmt.Printf("Counter: %d \n", counter)
}
```

### 未公开的标识符

把标识符命名为小写字母开头

### 公开的标识符

把标识符命名为大写字符的开头
