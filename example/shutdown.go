package main

import (
	"fmt"
	"github.com/qmstar0/shutdown"
)

func main() {
	shutdown.RegisterTasks(func() {
		fmt.Println("run task before exitting")
	})
	shutdown.WaitCtrlC()
}
