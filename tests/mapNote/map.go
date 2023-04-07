package mapNote

import "fmt"

func MapTesting() {
	fmt.Println("map 测试------------")
	//var mapTest map[int]string
	var a = make(map[int]string, 10)
	var b = make(map[int]string)
	c := map[int]string{
		23: "hello",
	}

	a[23] = "zhaoSI"
	a[45] = "sys_darwin"
	delete(c, 23)
	fmt.Println(a, "打印 MAP", b, c)

	// 查找
	value, flag := a[23]
	fmt.Println(value, flag, "查找打印")

	d := make(map[string]map[int]string)
	d["class"] = make(map[int]string)
	d["class"][1234] = "姗姗"

	fmt.Println("------->", d)
}
