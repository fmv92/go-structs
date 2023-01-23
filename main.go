package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact contactInfo
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
		contact: contactInfo{
			email: "faye@gmail.com",
			zip: 31000, 
		},
	}
	var spike person
	spike.firstName = "Spike"
	spike.lastName = "Spiegel"
	spike.contact.email = "bebop@gmail.com"
	spike.contact.zip = 31150

	fmt.Printf("%+v \n", spike)
	fmt.Printf("%+v", faye)

}