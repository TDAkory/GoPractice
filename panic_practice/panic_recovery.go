package panic_practice

import "fmt"

func A() {
	defer func() {
		fmt.Println("A.Func() 1")
	}()

	defer func() {
		fmt.Println("A.Func() 2")
		panic("abort in defer function")
	}()

	B()
	fmt.Println("should not be here")
}

func B() {
	defer func() {
		fmt.Println("defer func in B")
	}()
	panic("panic in B")
}

func C() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	A()
}
