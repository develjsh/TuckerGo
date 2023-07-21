// package는 코드를 묶는 단위입니다.
// 그래서 go code의 시작은 package로 시작해야합니다.
// package 이름은 어떤거든 적어도 되는데 main은 특별한 의미를 가지고 있습니다.
// main은 프로그램 시작점을 포함하는 패키지라는 의미입니다.
package main

import "fmt"

// 함수에서 main은 프로그램 시작점입니다.
func main() {
	fmt.Println("Hello World")
}
