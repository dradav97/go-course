package mypackage

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year int
}

// printMessage print a message
func PrintMessage(text string){
	fmt.Println(text)
}