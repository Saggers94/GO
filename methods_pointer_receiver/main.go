package main

import "fmt"

type car struct {
	brand string
	price int
}

type distance *int

//This is not permitted in go as you cannot have a method with a pointer receiver on a type
// that itself is a pointer
// func (d *distance) m1() {
// 	fmt.Println("just a message")
// }

// so we can only create a method on the value type not on the pointer type

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

// Here "(c car)" is the receiver of the method
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	// Here the parentheses are mandatory without them we will get an error
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Toyota", 30000)
	fmt.Println(myCar) // no change because the type is struct and string and int, which are passed by value

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar) // no change due to the same pass by value mechanism

	// The method change the value that pointer points to
	// (&myCar).changeCar2("BMW", 50000)
	// The go helps us here if the receiver is a pointer to a struct than the go will implicitly
	// convert the value to a pointer so, we don't need to use &myCar
	// we can just do the operation normally
	myCar.changeCar2("BMW", 50000)
	fmt.Println(myCar) // this will change because we are passing the pointer to the struct

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	// valid ways
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)

	// if a method takes a pointer receiver, it is good to convert all of the methods to pointer recerivers
	// it doesn't matter if they have to change the original value or not

	//1.  Use pointer receiver when you want to modify the receiver(changing the values) or
	//2.  when you want to avoid copying the large amount of data that is automatically copy when passing the value

	// if it's neither of the above two use value receiver
}
