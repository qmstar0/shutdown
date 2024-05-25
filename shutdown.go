package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	tasks  = make([]func(), 0)
	lock   sync.Mutex
	wg     sync.WaitGroup
	runOne = &sync.Once{}
	downCh = make(chan os.Signal, 1)
)

func init() {
	wg.Add(1)
	go func() {
		signal.Notify(downCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-downCh
		Exit(0)
	}()
}

func RegisterTasks(fn ...func()) {
	lock.Lock()
	defer lock.Unlock()
	tasks = append(tasks, fn...)
}

func WaitCtrlC() {
	wg.Wait()
}

func Exit(code int) {
	runOne.Do(func() {
		defer wg.Done()
		close(downCh)
		for i := range tasks {
			tasks[i]()
		}
		fmt.Println("\033[1mExit\033[m")
	})
	os.Exit(code)
}
