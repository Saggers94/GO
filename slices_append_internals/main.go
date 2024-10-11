package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// When we append to a slice that's backing array is full than go will create a new backing array
	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Appending another element will make the backing array capacity to four
	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 5)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Whenever we append an element to a specific index, go will replace the previous value with the new one
	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// Whenever backing array reaches its capacity in the next append operation,
	// go will create a new backing array with a length that's double than the current capacity
	// That's how slices do it

	letters := []string{"A", "B", "C", "D", "E", "F"}

	letters = append(letters[:1], "x", "y")
	fmt.Printf("Length: %d, Capacity: %d\n", len(letters), cap(letters)) // length is 3, capcity is 6

	// fmt.Println(letters[4]) error we cannot access after its length
	fmt.Println(letters[3:6]) // This would work because slices see the whole backing array

	// Unless and until the capcity is reached there will be no new backing array because of which
	// the data that was stored in the backing array would still be accessible, not by indexing
	// but by slicing
}
