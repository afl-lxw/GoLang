package DaisyChain

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		ch <- true // 到时间啦！
	}()
	return ch
}

func TimeAfterMain() {
	timeout := timer(time.Second) // 定时1s
	for {
		select {
		case <-timeout:
			fmt.Println("----already 1s!") // 到时间
			return                         //结束程序
		}
	}
}
