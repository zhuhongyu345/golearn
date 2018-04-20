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
	// 赋值 大写字母的 可以赋值 小写字母的不能,而且如果是指针，反射成功，如果是值，失败

	vp.Elem().FieldByName("Age").SetInt(28)
	vp.Elem().FieldByName("Name").SetString("zhu")
	fmt.Println(stu)



}

type stu struct {
	Name string
	Age  int
}

func (s stu) Dos() {
	fmt.Println("name", s.Name, "age", s.Age)
}

func Hello() {
	fmt.Println("hello")
}
