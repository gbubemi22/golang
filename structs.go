package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func getInput( prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	imput, err := r.ReadString('\n')
	return strings.TrimSpace(imput), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
     name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name) 
	fmt.Println("created the bill - ", b.name)
	return b 
} 

func main() {

	mybill := createBill()
	// mybill.addItem("onion soup", 4.50)
	// mybill.addItem("veg pie", 8.95)
	// mybill.addItem("toffee pudding", 4.95)
	// mybill.addItem("coffee", 3.25)

	fmt.Println(mybill)

}
