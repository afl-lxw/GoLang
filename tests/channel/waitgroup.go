package channel

import "fmt"

func Channel_main() {
	//var wg sync.waitgroup
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d ", i)
	//}
	//wg.Add(1)
	//
	//go func() {
	//	for i := 0; i < ; i++ {
	//		fmt.Printf("%d ", i)
	//	}
	//	defer wg.Done()
	//}
	//wg.Wait()
	//ChannelMake()
	//RangeLoop()
	TineMain()
}

func RangeLoop() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关闭信道 禁止流入 只可流出
	for v := range ch {
		fmt.Println(v)
		if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
			break
		}
	}
}

func ChannelMake() {
	//channel := make(chan int)
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("Ping!")

	fmt.Println(<-messages) // 取消息

	//messages <- "photoShop"
	//<-messages
}
