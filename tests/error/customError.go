package errors

import (
	"errors"
	"fmt"
)

func Init() {
	err := test()
	if err != nil {
		fmt.Printf("自定义错误%T", err)
		panic(err) // 执行中断
		fmt.Println("查看执行中段")
	}
}

func test() (err error) {
	num := 10
	num2 := 10
	if num2 == 0 {
		return errors.New("除数不能为0~~~~~")
	} else {
		result := num / num2
		fmt.Printf("执行成功 %T", result)
		return nil
	}
}
