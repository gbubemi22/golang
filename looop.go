package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
    fmt.Println("value of x is:", x)
    x++		
	}

// 
  names := []string{"mario", "luigi", "yoshi", "peach"}
    

//   for i := 0; i < len(names); i++ {
// 	fmt.Println(names[i])
//   }

// for index, value := range names{
// 	fmt.Printf("the position at index %v is %v \n", index, value)
// }



for _, value := range names{
	fmt.Printf("the position at is %v \n", value)
}

fmt.Println(names)
 
}