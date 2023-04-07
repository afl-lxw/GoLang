package class

import "fmt"

type Sleeper interface {
	Sleep()
}
type Eaterange interface {
	Eat(foodName string)
}
type LazyAnimal interface {
	Sleeper
	Eaterange
}
type Dog struct {
	Name string
}

func (this *Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping", this.Name)
}

type Cat struct {
	Name string
}

func (this Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping\n", this.Name)
}
func AnimalSleeper(this Sleeper) {
	this.Sleep()
}
func MainInter() {
	var s Sleeper
	dog := Dog{Name: "xiaoBai"}
	cat := Cat{Name: "helloKitty "}
	s = &dog
	AnimalSleeper(s)
	s = &cat
	AnimalSleeper(s)

	sleepList := []Sleeper{Cat{Name: "Kitty"}, Cat{Name: "---"}}
	for _, s := range sleepList {
		s.Sleep()
	}
}
