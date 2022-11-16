package main

import (
	"fmt"
	"strings"
)

// func sayGreeting(n string){
// 	fmt.Printf("Good morning %v \n", n)
// }  

// func sayBye(n string){
// 	fmt.Printf("GoodBye %v \n", n)
// } 

// func cycleNames(n []string, f func(string)){
// for _, v := range n {
// 	f(v)
// }	
// }

// func circleArea(r float64) float64{
// 	return math.Pi * r * r
// }

func getInitials(n string) (string, string) {
 s := strings.ToUpper(n)
 names := strings.Split(s, "")

 var initials []string
 for _, v := range names {
	initials = append(initials, v[:1])
 }
 if len(initials) >1 {
	return initials[0], initials[1]
 }

 return initials[0], "_"
}

func main () {
	
fn1, sn1 := getInitials("gbubemi work" )

fmt.Println(fn1, sn1)



fn2, sn2 := getInitials("Oritse code" )

fmt.Println(fn2, sn2)




fn3, sn3 := getInitials("Zoom " )

fmt.Println(fn3, sn3)

	//len append()
// sayGreeting("Gbubemi")
// sayGreeting("Oritsegbubemi")
// sayBye("Oritse")

// cycleNames([]string{"bewo", "busayo", "theodore"}, sayGreeting)
// cycleNames([]string{"bewo", "busayo", "theodore"}, sayBye)

// a1 := circleArea(10.5)
// a2 := circleArea(15)

// fmt.Println(a1, a2)
// fmt.Printf("cirle 1 is %0.3f and cirle 2 is %0.3f", a1, a2)



}