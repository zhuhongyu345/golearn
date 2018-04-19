package grammar

import (
	"testing"
	"fmt"
)

type Human struct {
	name string
	age  int
}
type Stu struct {
	Human
	school string
}

func (h Human) SayHai() {
	fmt.Println("say hai human")
}
func (s Stu) SayHai() {
	fmt.Println("say hai stu")
}
func TestInterface(t *testing.T) {
	fmt.Println("test interface")
	h := Stu{Human{"li", 22}, "bjd"}
	h.SayHai()

}
