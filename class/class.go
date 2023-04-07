package class

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	Schema string
}

func Main() {
	var t1 Teacher
	fmt.Println(t1)
	t1.Name = "ma"
	t1.Age = 34
	fmt.Println("class对象集合", t1)
}

func MainTest() {
	var T Teacher = Teacher{"刘孝文", 31, "山东"}
	fmt.Println(T, "------")
}
