package grammar

import (
	"testing"
	"fmt"
	"reflect"
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
func (s Stu) String() string {
	return fmt.Sprint(s.Human, "-", s.school)
}

/**
interface{}相当于泛型T
 */

type class interface {
}

func TestInterface(t *testing.T) {
	fmt.Println("test interface")
	h := Stu{Human{"li", 22}, "bjd"}
	h.SayHai()
	fmt.Println(h)

	l := make([]class, 3)
	l[0] = "str"
	l[1] = int32(4)
	l[2] = h
	fmt.Println(l)

	for i, class := range l {
		if i2, ok := class.(int); ok {
			fmt.Println(i, i2)
		} else if i3, ok := class.(string); ok {
			fmt.Println(i, i3)
		} else if i4, ok := class.(Stu); ok {
			fmt.Println(i, i4)
		} else {
			fmt.Println(reflect.TypeOf(class))
			fmt.Printf("%T\n", class)
		}

		switch v := class.(type) {
		case int32:
			fmt.Println(i, class, reflect.TypeOf(v))
		default:
			fmt.Println("s unknow",reflect.TypeOf(v))
		}
	}

}
