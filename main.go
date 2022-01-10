package main

import "fmt"

func main() {

	//variables
	var number uint16 = 250
	var language string = "Go-lang"
	fmt.Println("Hello Becky,welcome to Go-lang!")
	fmt.Println(number)
	fmt.Println(language)


	// advanced variable naming,for loops and string formatting, continue and break statement
	m := 51
	for m<= 51 {
		fmt.Printf("M is %v years old\n", m)
		break
	}

	z := 1
	for z<=5{
		fmt.Println(z)
		z = z+1
	}

	for h := 0; h<=10; h++ {
		if h%2 == 0{
			continue
		}
		fmt.Println(h)
	}
}