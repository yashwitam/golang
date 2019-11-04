package main

import (
	"fmt"
	"os"
)

func main() {
	var a = 10
	if(a>5){
		fmt.Println("a is greater than 5")
	}else{
		fmt.Println("a is less than 5")
	}

	//assign b a value and use value of b to evalaute
	if b := a/2; b>5{
		fmt.Println("b greater than 5", a,b)
	}else{
		fmt.Println("b less than 5", a,b)
	}

	//for statement: assignment, condition, increment
	for i := 0; i<10; i++{
		fmt.Println(i)
	}


	for i := 0; i<10; i++{
		if(i == 7){
			fmt.Println("i is : ", i)
		}
	}

	for i := 0; i<10; i++{
		if(i == 7){
			fmt.Println("breaking ")
			break //can also use continue
		}else{
			fmt.Println(i)
		}
	}

	s1:= "Hello, World"
	for k, v:= range s1{
		fmt.Println(k,v, string(v)) //v is rune
	}


	//switch statements
	word :=os.Args[1]
	if word == "hello"{
		fmt.Println("printing hello")
	}else if word == "goodbye"{
		fmt.Println("printing goodbye")
	}else{
		fmt.Println("landing into no zone")
	}





}
