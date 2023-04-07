package main

import (
	errors "Golang/error"
	"Golang/mysql"
	"fmt"
)

func main() {
	errors.Init()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("错误捕获 %T", err)
			panic(err)
		}
	}()
	//fmt.Println("你好")
	//note.GoSecFun()
	//noteA.SaySA()
	//noteA.Init()
	//errors.Init()
	//arrays.Array()
	//array := make([]int, 0)
	//array = append(array, 1, 2, 3, 4, 5)
	//result := slices.Slices(array)
	//fmt.Println(result, "<-----")
	//slices.MakeSlice()
	//mapNote.MapTesting()
	//class.Main()
	//class.MainTest()
	//class.Mains()
	//class.Init("duck", 456)
	//class.MainInter()
	//channel.Main()
	//channel.Channel_main()
	//generator.GeneratorMain()
	//multiplexChannel.MultiplexMain()
	//multiplexChannel.SelectMain()
	//DaisyChain.DaisyMain()
	//DaisyChain.TimeAfterMain()
	//toDo.TodoMain()
	//toDo.UnBufMain()
	mysql.Mysql_start()
}
