package grammar

import (
	"testing"
	"fmt"
	"strconv"
)

func TestPro(t *testing.T) {
	fmt.Println("test begin")

	i := 0
H:
	if i < 3 {
		i++
		fmt.Println("here" + strconv.Itoa(i))
		goto H
	}

	for ii := 1; ii < 10; {
		ii += ii
		fmt.Println(ii)
	}
	for i < 5 {
		i++
	}
	fmt.Println(i)

	user := map[string]string{"name": "zh", "age": "33"}
	user["job"] = "test"
	for k, v := range user {
		fmt.Println(k, v)
	}

	k := "xu"
	switch k {
	case "xu":
		fmt.Println("xu")
		k = "xu1"
		fallthrough
	case "xu1":
		fmt.Println("xu")
	case "zhu":
		fmt.Println("zhu")
	default:
		fmt.Println("unknown")

	}

	manyParam("zhu", 4, 3, 2)

	add := 2
	addone(&add)
	fmt.Println(add)

	s := "aa"
	stringadd(s)
	fmt.Println(s)

	defer fmt.Println(222)
	m := make(map[string]int)
	mapadd(m)
	fmt.Println(m)

	defer fmt.Print(1111)

	sli := make([]int, 33, 33)
	fmt.Println(sli)
	sliceadd(sli)
	fmt.Println(sli)

	testPanic()

	type Person struct {
		name string
		age  int
	}
	p := Person{"22", 33}
	fmt.Println(p)
	fmt.Println(p.age)

	type WPerson struct {
		*Person
		color string
		string
	}
	person2 := WPerson{&p, "ss", "www"}
	fmt.Println(person2)
	fmt.Println(person2.Person.age)
	person2.string = "ppp"
	person2.Person.age = 44
	fmt.Println(person2)
	fmt.Println(p)

	type Human struct {
		phone string
	}
	type Employee struct {
		Human
		phone string
	}
	Bob := Employee{Human{"777-444-XXXX"}, "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}

func testPanic() {
	defer func() {
		i := recover()
		fmt.Println(i)
	}()
	panic("test error")
}

func sliceadd(s []int) {
	//ints := append(s, 1,2,3,4)
	//fmt.Println(ints)
	s[9] = 1
}

func manyParam(a string, bs ... int) {
	fmt.Println(a)
	for i, p := range bs {
		fmt.Println(i, p)
	}
}

func addone(a *int) {
	*a += 1
}

func stringadd(s string) {
	s = "a"
}
func stringaddp(s *string) {
	*s = *s + *s
}

func mapadd(m map[string]int) {
	m["aaa"] = 0
}
