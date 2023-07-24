package main

import "fmt"

const Y int = 3

func main() {
	// X:= 5
	// a := [X]int{1,2,3,4,5} // error 발생합니다. X는 변수이기 떄문에 error가 발생합니다.

	b := [Y]int{1,2,3}

	var c [10]int

	fmt.Println(b, c)
}