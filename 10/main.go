package main

import (
	room "ex10/room"
	"fmt"
)

func main(){
	fmt.Println("Hello, Gopher!")
	num := room.RoomLights()
	fmt.Println("number = ", num)
}