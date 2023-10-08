package test

import "fmt"

func TestRegisterTplRouter() {
	x := 1
	y := 2
	defer calcTest("AA", x, calcTest("A", x, y))
	x = 10
	defer calcTest("BB", x, calcTest("B", x, y))
	y = 20
}

func calcTest(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
