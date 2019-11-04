package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	passByReferenceExample(b)
	fmt.Println("passByReferenceExample: ",*b)
	fmt.Println("passByReferenceExample: ",a)

	passByValueExample(b)
	fmt.Println("passByValueExample: ",*b)
	fmt.Println("passByValueExample: ",a)

	nilPointerExample()

}


func passByReferenceExample(b *int){

	fmt.Println(*b)
	*b = 20
	fmt.Println(*b)
}

func passByValueExample(b *int){ //value is the address of memory location for a pointer

	fmt.Println(*b)
	b = new(int)
	*b = 30
	fmt.Println(*b)

}


func nilPointerExample(){
	var b *int

	fmt.Println(b)
	//fmt.Println(*b) ---> will give nil pointer derefence


	c := new(int) //poniter to the type
	fmt.Println(c) // prints address
	fmt.Println(*c) // prints 0
}