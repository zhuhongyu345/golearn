package grammar

import (
	"testing"
	"time"
	"fmt"
)
const i = iota
const j = iota
const k = iota

const (
	ii = iota
	jj = iota
	uu = 5
	kk = iota
)

func HelloWorld() {

	fmt.Printf("Hello, world %v你好，世界 %vκαλημ ́ρα κóσμ %vこんにちは世界%v", "\n", "\n", "\n", "\n")
	time.Sleep(time.Microsecond)

	s := "hello"
	c := []byte(s)
	fmt.Println(s + s)
	fmt.Println(string(c[0]))
	fmt.Println(string(c[1]))

	m := `hello
world`
	fmt.Println(m)

	fmt.Println(i, j, k)
	fmt.Println(ii, jj, kk)

	ss := []int{1, 2, 3, 4, 5}
	fmt.Println(ss)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr[0] = 10

	ss = arr[1:7]
	ss[2] = 10
	fmt.Println(ss)
	fmt.Println(arr)
	fmt.Println(len(ss), cap(ss))

	i := append(ss, 122)

	fmt.Println(i)
	fmt.Println(len(ss), cap(ss), ss)

	smap := map[string]string{"aa": "192", "bb": "122", "cc": "10"}
	fmt.Println(smap)
	i2, ok := smap["aaa"]
	fmt.Println(i2, ok)
	i2, ok = smap["aa"]
	fmt.Println(i2, ok)
	delete(smap,"aa")
	delete(smap,"aaa")
	fmt.Println(smap)

}
func TestHelloWorld(t *testing.T)  {
	 HelloWorld()
	time.Sleep(time.Microsecond*5)

}