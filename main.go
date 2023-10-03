package main

import (
	"POO_go/classes"
	"fmt"
)

func main() {
	address := classes.Address{
		Street: "123 Main St",
		City:   "Cityville",
		State:  "Stateville",
	}
	person := classes.NewPerson("Alice", 30, []string{"Reading", "Painting"}, address)

	person.DisplayBasicInfo()
	person.DisplayHobbies()
	person.DisplayAddress()
	fmt.Printf("Age in days: %d\n", person.CalculateAgeInDays())
}
