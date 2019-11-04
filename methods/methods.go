package main

import "fmt"

type Person struct{
	name string
	age int
}


func main() {

	fmt.Println(String())

	p := Person{"Yash", 25}

	fmt.Println(p.methodExample()) //notice you have to intitialize the struct to invoke the methodExample method on person
	p.methodByReferenceExample()
	fmt.Println(p.methodExample())
}

//usual function declaration
func String() string{
	return "Hello World!"
}

//invoking method using value receiver example
//if you do not want to modify values in the struct use value receiver
func (person Person) methodExample() string{ //person receivable, copy of person is passed
	//return "Hello " + person.name + "!"
	return fmt.Sprintf("Hello %s! Your age is:%d", person.name, person.age)
}

//invoking method using reference receiver example
//if you want to modify values in the struct use reference receiver
func(person *Person) methodByReferenceExample(){
	person.age = person.age + 5
}
