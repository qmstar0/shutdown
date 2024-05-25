# shutdown

shutdown 包提供了一种优雅的方式来处理 Go 应用程序的关闭信号，执行注册的清理任务，并确保在应用程序退出之前完成所有注册的清理任务。

## 安装

在您的项目中使用 go get 来安装此包：

```
go get github.com/qmstar0/shutdown
```

## 示例
```
package main

import (
	"fmt"
	"github.com/qmstar0/shutdown"
)

func main() {
	// 注册清理任务
	shutdown.RegisterTasks(func() {
		fmt.Println("执行清理任务1")
	}, func() {
		fmt.Println("执行清理任务2")
	})

	// 模拟其他工作
	fmt.Println("程序正在运行...")

	// 等待退出
	shutdown.WaitCtrlC()
}

```
在这个示例中，程序运行时会注册两个清理任务，并在接收到关闭信号后执行这些任务，然后地退出程序。
