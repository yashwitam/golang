package main

import "fmt"

func main() {
	var s1 = "Hello String!"
	fmt.Println(s1)

	s2 := "Another way of declaring string, using type inference operator"
	fmt.Println(s2)

	var s3 string = "Yet another way of declaring string"
	fmt.Println(s3)

	s4 := "Hello"
	s5 := "World"
	fmt.Println(s4+" "+s5)

	s6 := "Hello World"
	s7 := s6[0]
	fmt.Println(s7) //prints 72
	fmt.Println(string(s7)) //prints character 'H'

	//substrings example
	s8 := "Hello, World"
	s9 := s8[0:5] //prints Hello
	s10 := s8[:5] //prints Hello
	fmt.Println(s9, s10)
	fmt.Println(s8[7:12], s8[7:]) //prints World World

	//lnegth of string
	s11 := "Hello, World"
	fmt.Println(len(s11)) //prints 12

	s12 := "Hello, "
	var r rune = 127757
	fmt.Println(s12 + string(r))
}
