package grammar

import (
	"testing"
	"fmt"
	"reflect"
)

func TestRef(t *testing.T) {

	a := float32(3.0)
	of := reflect.TypeOf(a)
	fmt.Println(of)
	o := reflect.ValueOf(a)
	//o.SetFloat(3.2)
	oo := reflect.ValueOf(&a)

	elem := oo.Elem()
	elem.SetFloat(3.4)

	fmt.Println(o)
	fmt.Println(o.Kind())

	stu := stu{"xu", 23}
	v := reflect.ValueOf(stu)
	fmt.Println(v)
	ty := reflect.TypeOf(stu)
	fmt.Println(ty)

	hl := reflect.ValueOf(Hello)
	hl.Call(nil)
	fmt.Println(hl.Kind() == reflect.Func)

	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf("zhu")
	params[1] = reflect.ValueOf(28)
	v.MethodByName("Dos").Call(nil)
	fmt.Println(v.Method(0))
	v.Method(0).Call(nil)


	vp := reflect.ValueOf(&stu)
	ele := vp.Elem()
	fmt.Println(ele)
	typ := ele.Type()
	fmt.Println(typ)
	field := typ.NumField()
	fmt.Println(field)
	for i := 0; i < field; i++ {
		fmt.Println(typ.Field(i).Name)
	}
	for i := 0; i < typ.NumMethod(); i++ {
		fmt.Println(typ.Method(i).Name)
	}
	fmt.Println(vp.Elem().Field(0))
	//todo  赋值失败

	fmt.Println(stu)



}

type stu struct {
	name string
	age  int
}

func (s stu) Dos() {
	fmt.Println("name", s.name, "age", s.age)
}

func Hello() {
	fmt.Println("hello")
}
