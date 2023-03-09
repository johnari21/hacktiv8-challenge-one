package main

import "fmt"

func main() {

	var (
		i             int     = 21
		j             bool    = true
		russiaChar    rune    = 'Я'
		russiaUnicode string  = "\u042F"
		x             float64 = 123.456
	)

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	fmt.Println(j)
	fmt.Printf("%b\n", i)
	fmt.Printf("%v\n", russiaUnicode)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%U\n", russiaChar)
	fmt.Printf("%f \n", x)
	fmt.Printf("%E \n", x)
}
