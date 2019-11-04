package main

import (
	"fmt"
	"strings"
	"math"
	"time"
	"github.com/ymishra/leftpad"
)

func main() {
	//str1 := "This is a string"
	s2 := strings.Compare("b","bc")
	fmt.Println(s2)

	t := time.Now()
	fmt.Println(t)

	m := math.Max(2,4)
	fmt.Println(m)

	str := leftpad.Format("Append three zeros to left and uppercase this string")
	fmt.Println(str)
}
