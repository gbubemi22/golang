package main


import (
	"fmt"
)

func updateName(x *string){
	*x = "wedge"
}


func main () {
	name:= "tifa"


	// updateName(name)

	//fmt.Println("memory address of name is:  ", &name)

	m := &name

	fmt.Println(name)
	

	updateName(m)
	fmt.Println(name)

	
}