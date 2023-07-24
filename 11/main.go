package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에이컨을 켠다.")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다.")
	} else if temp <= 18 {
		fmt.Println("밖에 나간다.")
	} else {
		fmt.Println("")
	}

	if filename, success := initFunc(); success {
		fmt.Println("Upload success", filename)
	} else {
		fmt.Println("Upload fail", filename)
	}

}

func initFunc() (string, bool){
	return "파일이름", true
}