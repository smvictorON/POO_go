package classes

import (
	"fmt"
)

// Person é uma struct que representa uma pessoa.
type Person struct {
	Name    string
	Age     int
	Hobbies []string
	Address Address
}

// NewPerson cria uma nova instância de Person.
func NewPerson(name string, age int, hobbies []string, address Address) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Hobbies: hobbies,
		Address: address,
	}
}

// DisplayBasicInfo exibe informações básicas da pessoa.
func (p *Person) DisplayBasicInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// DisplayHobbies exibe os hobbies da pessoa.
func (p *Person) DisplayHobbies() {
	fmt.Printf("Hobbies: %s\n", p.Hobbies)
}

// DisplayAddress exibe o endereço da pessoa.
func (p *Person) DisplayAddress() {
	fmt.Printf("Address: %s, %s, %s\n", p.Address.Street, p.Address.City, p.Address.State)
}

// CalculateAgeInDays calcula a idade da pessoa em dias.
func (p *Person) CalculateAgeInDays() int {
	return p.Age * 365 // Aproximação simples de dias em um ano
}
