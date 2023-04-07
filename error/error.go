package errors

import "fmt"

func Error() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(" 错误捕获 %T", err)
		}
	}()
}
