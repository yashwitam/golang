package main

import "fmt"

func addOne(a int) int{
	return a + 1;
}

func addTwo(b int) int{
	return b+2;
}

//example passing functions
func printFunction(a int, f func(int) int){
	fmt.Println(f(a))
}

//example passing functions
func makeAdder(b int) func(int) int{
	return func(a int) int{
		return a+b
	}
}

func makeDoubler(f func(int) int) func(int) int{
	return func(a int) int{
		b:=f(a)
		return b*2
	}
}

func main() {
	fmt.Println(addOne(1))
	myOne := addOne           //assign function to a variable
	fmt.Println(myOne(2))

	myTwo := func(b int) int{ // myTwo is just another variable to which an anonymous function is assigned. anonymous functions can be defined inside of main
		return b+2
	}

	fmt.Println(myTwo((2)))

	printFunction(10, addOne) //prints 11 i.e you can pass around functions like data variables
	printFunction(11, addTwo) //prints 13


	//closure example
	d := 1
	closureExample := func(a int) {
		d = d + a
	}

	closureExample(1)
	fmt.Println(d) // prints value of d as 2

	//example passing functions
	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(1))
	fmt.Println(addTwo(1))

	//example passing cuntion returning function
	addOne1 := makeAdder(1)
	doubleAddOne := makeDoubler(addOne1)

	fmt.Println("printing addone1: ",addOne1(1))
	fmt.Println("printing doubleAddOne: ",doubleAddOne(1))



}

