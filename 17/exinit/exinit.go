package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f() // c 값을 초기화 하기 위해 f() 함수를 실행하고 b 값을 초기화 하기 위해 f() 험수를 실행하기 때문에 2번 f() 함수는 2번 실행하게 됩니다.
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init funciton", d)
}

func f() int {
	d++
	fmt.Println("f() d: ", d)
	return d
}

func PrintD() {
	fmt.Println("PrintD d: ", d)
}