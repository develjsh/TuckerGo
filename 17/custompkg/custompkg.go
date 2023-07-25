package custompkg

import (
	"fmt"
	"tuckergo/usepkg/exinit"
)

func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD()
}

func PrintCustom2() {
	fmt.Println("This is custom package2222!")
}