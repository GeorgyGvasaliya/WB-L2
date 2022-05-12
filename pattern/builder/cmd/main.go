package main

import (
	"L2/pattern/builder/pkg"
	"fmt"
)

func main() {
	normalBuilder := pkg.GetBuilder("wooden")

	director := pkg.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

}
