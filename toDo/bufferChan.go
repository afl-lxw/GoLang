package toDo

import (
	"fmt"
	"time"
)

func UnBufMain() {
	unbufChan := make(chan int)
	go func() {
		fmt.Printf("Sleep a second ...\n")
		time.Sleep(time.Second)
		num := <-unbufChan
		fmt.Printf("Received a integer %d.\n", num)

	}()
	num := 1
	fmt.Printf("Send integer %d ...\n", num)
	unbufChan <- num
	fmt.Printf("Done")
}
