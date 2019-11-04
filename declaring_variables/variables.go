package main

import "fmt"


func main() {
	var i int = 10 //explicit declaration of variable using data type
	fmt.Println(i)

	var j = 12 //declaration of variable with inferred data type
	fmt.Println(j)

	k := 14 //implicit inference of data type, using := operator, without using var keyword or datatype
	fmt.Println(k)

	var t int
	fmt.Println(t)
	t = 20
	fmt.Println(t)


}
