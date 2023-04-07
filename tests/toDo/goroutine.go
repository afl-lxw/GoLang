package toDo

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func TodoMain() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv2, sigs2...)
	// 接着我们用两个for循环来接收消息.考虑到主程序会退出，所以用waitgroup等待完成后退出.

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1%v", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv1%v", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
	}()

	wg.Wait()
}
