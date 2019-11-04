package main

import ("fmt"
	"os"
)

func main() {
	word := os.Args[1]
	switch word{
	case "hi":
		fmt.Println("hi")
		fallthrough
	case "hello":
		fmt.Println("hello")
	case "greetings", "bye":
		fmt.Println("greetings or bye")
	default:
		fmt.Println("default case")
	}



	word2 := os.Args[2]
	c := "crackerjack"
	switch l := len(word2); {
	case word2 == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word2 == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word2 == c:
		fmt.Println("This word is either", c, "or 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long.")
	}

	switchexample("howdie")
}

func switchexample(word string){
	switch word{
	case "hello":
		fmt.Println("hello")
	case "howdie":
		fmt.Println("howdie")
	}
}