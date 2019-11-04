package main

import "fmt"

type Person struct{
	age int
	name string
}


func main() {

	var person Person
	fmt.Println(person)

	person = Person{10, "yash"}
	fmt.Println(person)

	person = Person{age:20}
	fmt.Println(person)

	person = Person{name:"amol"}
	fmt.Println(person)

	person = Person{10, "amol"}
	fmt.Println(person)


	var person1 *Person //struct pointer
	person1 =&person
	person1.name = "yash"
	fmt.Println(*person1)
	fmt.Println(person)
}
