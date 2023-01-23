package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func main() {

	//spike := person{"Spike", "Spiegel"}
	faye := person{
		firstName: "Faye",
		lastName:  "Valentine",
		contactInfo: contactInfo{
			email: "faye@gmail.com",
			zip: 31000, 
		},
	}
	
	var spike person
	spike.firstName = "Spike"
	spike.lastName = "Spiegel"
	spike.updateName("Jet")
	spike.contactInfo.email = "bebop@gmail.com"
	spike.contactInfo.zip = 31150

	spike.print()
	faye.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}