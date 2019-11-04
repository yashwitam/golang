package main

import "fmt"

func main(){

	in := make (chan string)
	out := make (chan string)

	go func() {
		name := <-in
		out <- fmt.Sprintf("Hello, ", name)
	}()

	in <- "Yash"
	close(in)
	message := <-out
	fmt.Println(message)
}
