## Go Var Listener
监听某个变量的变化，当变量发生变化时，会触发回调函数。

### 使用方法
**具体使用方法**请查看[测试集](api_test.go)  
以下使用介绍可能因为API改变而不准确  
```go
package main

import (
	"fmt"
	gvl "git.lolli.tech/lollipopkit/go-var-listener"
	"time"
)

func main() {
	vv := 1
	// 新建监听
	v := gvl.NewVar(1)
	v.Listen(Callback[int]{
		fn: func(i int) {
			vv += 1
		},
		// 该回调函数的唯一ID
		name: "add1-onboth",
		// 类型：改变值时、读取时调用，或者两者都调用
		typ: gvl.OnBoth,
	})
	// 查看监听是否设置成功
	v.HaveListener("add1-onboth") // true
	// 修改变量
	v.Set(2)
	// 等待回调函数执行
	time.Sleep(time.Second)
	fmt.Println(vv) // 2
	// 读取
	v.Get() // 2
	time.Sleep(time.Second)
	fmt.Println(vv) // 3
	// 取消监听
	v.Unlisten("add1-onboth")
}

```