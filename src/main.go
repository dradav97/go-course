package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Fearri"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("hellow world")

}