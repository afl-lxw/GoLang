package slices

import "fmt"

func Slices(arrays []int) []int {
	arr := arrays[1:3]
	// cap 查看切片容量
	fmt.Println(arr, "切片")
	println(len(arr), "切片长度")
	return arr
}

func MakeSlice() {
	slice := make([]int, 4, 20)
	fmt.Println(slice, "================================")
	fmt.Println(len(slice), "切片的长度")
	fmt.Println(cap(slice), "切片的容量")
	slice[0] = 12
	slice[1] = 124
	slice[3] = 123
	slice[2] = 30
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println("循环打印==》", slice[i])
	}
	for m, n := range slice {
		fmt.Println("range循环打印==》", m, n)
	}
	fmt.Println("切片简写-----------------")

}
