package go_grammar

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Eat(food string) {
	fmt.Printf("%s is eatting an %s\n", p.Name, food)
}

type Student struct {
	Person
	StudentId int
}

func (s *Student) Eat(food string) {
	fmt.Printf("%s student is eatting a %s\n", s.Name, food)
}

func TestStruct(t *testing.T) {
	person := &Person{
		Name: "lpz",
		Age:  20,
	}
	person.Eat("apple")

	student := &Student{
		Person: Person{
			Name: "lz",
			Age:  25,
		},
		StudentId: 1,
	}

	student.Eat("banana")
}
