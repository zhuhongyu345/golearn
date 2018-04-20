package grammar

import (
	"testing"
	"fmt"
	"runtime"
)

func TestGoroutine(t *testing.T) {
	sumCun([]int{1, 2, 3, 4, 5, 6, 7, 8})

	fmt.Println()

	cap1 := 1
	cc := make(chan int, cap1)
	go fbnqSync(cap1, cc)
	for j := range cc {
		fmt.Println(j)
	}
	fmt.Println()

	ccc := make(chan int, cap1)
	fbnqDg(cap1, ccc)
	for j := range ccc {
		fmt.Println(j)
	}
	runtime.Gosched()

	cccc := make(chan int, 5)
	quit := make(chan int, 2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cccc)
		}
		quit <- 0
	}()

	go twoChan(cccc, quit)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOROOT())

	mkerr()
	fmt.Println("normal")
}

func mkerr() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("error")
}

func twoChan(c, quit chan int) {
	//runtime.Goexit()
	x, y := 1, 1
L2:
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case a := <-quit:
			fmt.Println(a)
			break L2
			//default:
			//	fmt.Println("default")
		}
	}
}

func fbnqDg(cap int, c chan int) {
	if cap == 1 {
		c <- 1
		close(c)
		return
	}
	c <- 1
	c <- 1
	fbnqDgAct(1, 1, cap, c)

}

func fbnqDgAct(a int, b int, cap int, c chan int) {
	if len(c) < cap {
		t := a + b
		c <- t
		fbnqDgAct(b, t, cap, c)
	} else {
		close(c)
	}
}

func fbnqSync(cap int, c chan int) {
	if cap == 1 {
		c <- 1
		close(c)
		return
	}
	a, b := 1, 1
	c <- a
	c <- b
	for i := 2; i < cap; i++ {
		t := a + b
		a, b = b, t
		c <- t
	}
	close(c)
}

func sumCun(data []int) {
	c := make(chan int, 3)
	go sum(data[:2*(len(data)/3)], c)
	go sum(data[len(data)/3:2*(len(data)/3)], c)
	go sum(data[0:len(data)/3], c)
	for len(c) < 3 {
	}
	close(c)
	for j := range c {
		fmt.Println(j)
	}
}

func sum(data []int, c chan int) {
	fmt.Println("sum begin")
	sum := 0
	for _, v := range data {
		sum += v
	}
	c <- sum
	fmt.Println("sum end")
}
