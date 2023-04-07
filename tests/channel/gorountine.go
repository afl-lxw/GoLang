package channel

import (
	"fmt"
	"runtime"
)

// 等待多gorountine的方案
var quit chan int = make(chan int)

func foo() {
	for i := 0; i < 10; i++ { //为了观察，跑多些
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func TineMain() {
	//runtime.GOMAXPROCS(2) // 最多使用2个核

	go foo()
	go foo()

	for i := 0; i < 2; i++ {
		<-quit
	}
}
