package main

import (
	"fmt"
)
/* 
//Functions used as multiple patameters to other function

type Mod func(s string) string 

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func doubler(s string) string {
	return s + s
}

func applyer(s string, mods ...Mod) string {
	for _, mod := range mods {
		s = mod(s)
	}

	return s
}

func main() {
	fmt.Println(applyer("Gay", toUpper, doubler))
}
*/


/* 
//Function can return fuction, so it can be used as variable 
func giveMeFunc() func(string) {
	return func(name string) {
		fmt.Println("Hello " + name + "!")
	}
}

func main() {
	greetFunc := giveMeFunc()
	greetFunc("Peder!")
}
*/

func incrementor() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main() {
	inc := incrementor()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
