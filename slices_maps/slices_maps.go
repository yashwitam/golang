package main

import "fmt"

func main() {
	s := make([]string, 0)
	fmt.Println(s)
	fmt.Println("length at line8 is: ", len(s))

	s = append(s,"hello")
	fmt.Println(s)
	fmt.Println("length at line12 is: ", len(s))

	s = append(s,"goodbye", "world")
	fmt.Println(s)
	fmt.Println("length at line12 is: ", len(s))


	c := make([]string, 2)
	c = append(c,"hi", "how", "are", "you")
	fmt.Println(c)
	fmt.Println("length of c at line22 is: ", len(c))


	d := []string{"another", "way", "to", "declare","slice"}
	fmt.Println(d)
	fmt.Println("length of d at line27 is: ", len(d))

	g := d[1:3]
	fmt.Println(g)
	fmt.Println("length of g at line31 is: ", len(g))

	g[0] = "changedfirstelement"
	fmt.Println(g)

	modeSlicePassByReferenceExample(g) //SLICES ARE REFERENCE TYPES
	fmt.Println(g)

	//iterate over slice
	for k, v := range g{
		fmt.Println(k,v)
	}


	m := map[int]string{
		1:"yash",
		2:"amol",
	}

	fmt.Println(m)
	fmt.Println(m[1])
	fmt.Println(m[2])
	fmt.Println("map length: ", len(m))

	modMap(m)
	fmt.Println(m)

	delete(m,1)
	fmt.Println(m)
	fmt.Println("map length: ", len(m))

	m[1] = "yash"
	m[3] = "vicky"

	//iterate over map
	for k, v := range m{
		fmt.Println("map elements", k,v)
	}
}

func modMap(t map[int]string){
	t[1] = "vicky"
}

func modeSlicePassByReferenceExample(s []string){
	s[0] = "modified"
}