package room

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room // &^ 비트 클리어
}

func IsLightOn(rooms, room uint8) bool {
	return rooms & room == rooms
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, uint8(MasterRoom)) {
		fmt.Println("MasterRoom 불을 켠다.")
	}
	if IsLightOn(rooms, uint8(LivingRoom)) {
		fmt.Println("LivingRoom 불을 켠다.")
	}
	if IsLightOn(rooms, uint8(BathRoom)) {
		fmt.Println("BathRoom 불을 켠다.")
	}
	if IsLightOn(rooms, uint8(SmallRoom)) {
		fmt.Println("SmallRoom 불을 켠다.")
	}
}

func RoomLights() int{
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	rooms = SetLight(rooms, BathRoom)

	rooms = ResetLight(rooms, uint8(SmallRoom))

	TurnLights(rooms)
	return 1
}