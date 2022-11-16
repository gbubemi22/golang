package main

import (
	"fmt"
)

func main () {
  menu := map[string]float64	{
	"soup": 4.99,
	"pie": 7.99,
	"salad": 6.99,
	"toffee padding": 3.99,
  }


  fmt.Println(menu)
  fmt.Println(menu["pie"])
  fmt.Println(menu["salad"])
  fmt.Println(menu["soup"])

  //looping maps

  for k, v := range menu{
	fmt.Println(k, "-", v)
  }


  // ints as key type

  phonebook := map[int]string{
	7030762738: "Gbubemi",
	8111631084: "Busayo",
	8030762738: "Theodore",

  }


  fmt.Println(phonebook)
  fmt.Println(phonebook[7030762738])
  fmt.Println(phonebook[8030762738])

  phonebook[8030762738] = "Eyituoyo"
  fmt.Println(phonebook)

  phonebook[8111631084] = "Oyindamola"
fmt.Println(phonebook)

}