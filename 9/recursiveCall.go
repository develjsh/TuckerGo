package main

import "fmt"

func main() {
	printNo(3) // 함수 호출
}

func printNo(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	printNo(n-1)	// 재귀 호출
	fmt.Println("After", n) // 재귀 호출 이후 출력
}