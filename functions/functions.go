package main

import "fmt"

func main() {
	addNumbers()
	addNumbers1(10, 20)
	fmt.Println(addNumbers2(30, 40))
	fmt.Println(swap(2,3))

	a,b := swap(10,20)
	fmt.Println(a, b)

	_,b = swap(40,60) //ignore the left side of return value after swapping
	fmt.Println(b)

	arrayexample([2]int{1,2}) //go is dumb it cannot identify two items in the int array, but instead needs it specified as [2]

}

func addNumbers(){
	fmt.Println(2 + 3)
}

func addNumbers1(a int, b int){ //does not support function overloading
	fmt.Println(a+b)
}

func addNumbers2(a int, b int) int{
	return a + b
}

func swap(a int, b int) (int, int){
	return b, a
}

func arrayexample(arr [2]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	fmt.Println(arr)
}