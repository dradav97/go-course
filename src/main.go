package main

import "fmt"

func normalFunction(message string){
	fmt.Println(message)
}

func threeArguments(a,b int, c string) {
	fmt.Println(a,b,c)
}

func returnValue (a int) int {
	return a * 2
}

func doubleReturn(a int) (c,d int) {
	return a, a*2
}

func main() {
	
}
