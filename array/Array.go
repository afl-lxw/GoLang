package arrays

import (
	"fmt"
)

func Array() {
	array := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(array))

	for key, value := range array {
		println(key, "------->", value)
	}
}
