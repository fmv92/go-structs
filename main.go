package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//spike := person{"Spike", "Spiegel"}
	// spike := person{
	// 	firstName: "Spike",
	// 	lastName:  "Spiegel"}
	var spike person
	spike.firstName = "Spike"
	spike.lastName = "Spiegel"

	fmt.Printf("%+v", spike)

}