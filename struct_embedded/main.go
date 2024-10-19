package main

import "fmt"

func main() {
	// nested or embedded struct
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact // This is an embedded struct
	}

	john := Employee{
		name:   "John Doe",
		salary: 4000,
		contactInfo: Contact{
			email:   "jd@company.com",
			address: "Street 20, London",
			phone:   0034,
		},
	}
	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)

	// struct is similiar to an object in javascript in which we can have nested objects
	// and its properties can be accessed using the dot operator
	// also we can modify the value with the equal operator too just like an object

	myContact := Contact{
		email:   "andrei@domain.com",
		phone:   23421323132,
		address: "Bucharest, Romania",
	}
	fmt.Println(myContact)
}
