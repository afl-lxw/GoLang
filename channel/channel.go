package channel

import (
	"fmt"
	"time"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("----------- \n")

}

func Main() {
	go loop() // 启动一个goroutine
	loop()
	loop()

	time.Sleep(time.Second) // 停顿一秒
}
