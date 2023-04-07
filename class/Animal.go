package class

type animal struct {
	Name string
	Age  int
}

type Tiger struct {
	animal
	Height int
	Space  string
}

func Init(name string, age int) *animal {
	return &animal{name, age}
}

func (this *animal) Say() string {
	this.Name = "element"
	println("this-->", this)
	return ""
}

func Mains() {
	Dog := animal{
		Name: "新衣",
		Age:  345,
	}
	Dog.Say()
}
