package main

import "fmt"

type House struct {
	Address string
	Size int
	Price float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구"
	house.Size = 28
	//house.Price = 10

	fmt.Println(house)
	fmt.Printf("주소:%s 사이즈:%d \n", house.Address, house.Size)
}