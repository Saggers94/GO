package main

import "fmt"

func main() {
	// Maps has a Maps header property that points to the underlying data structure in memory
	// so it saves only the reference to the underlying data structure
	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbours := friends

	friends["Dan"] = 50
	// Neighbours map has been modified because it has the same reference to friends
	fmt.Println(neighbours)

	// To clone a map we have to create a new map and than copy all the elements
	// from the original map to the new map via for loop
	people := make(map[string]int)
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18
	// People map has been modified but not the friends map
	fmt.Printf("%#v --------------------- %#v \n", people, friends)

	// This will delete the key and value from the friends map but not from the people map
	// as it is a different map
	delete(friends, "Dan")
	fmt.Printf("%#v ........... %#v", friends, people)

	// It is not necessary to check if the key exists or not before deleting it from the map
	// The program will run without an errror even if the key doesn't exist
	delete(people, "Andre")

}
