package main

import "fmt"

func main() {
	age := 50
	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println("adult age:", editAgeToAdultYears(agePointer))
}
func editAgeToAdultYears(age *int) int {
	return *age - 18
}
