package main

import "fmt"

func main() {
	// var은 variable이라는 뜻입니다.
	// a는 변수명입니다.
	// int는 타입입니다.
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
}
