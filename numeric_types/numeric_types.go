package main

import "fmt"

func main() {
	//types int8, int16, float16, float32, bool, byte
	var  i = 8
	var  j = 3.5
	fmt.Println((float64(i) + j))

	fmt.Println(i + int(j))
}
