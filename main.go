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
	
	// calling functions
	a:= 20 
	b:= 8
	response := addition(a, b)
	fmt.Printf("If you add %v + %v you get %v \n", a, b, response)

	sub_response := substration(a, b)
	fmt.Printf("Substracting %v from %v the answer is %v \n", b, a, sub_response)

	mult_response := multiplication(a, b)
	fmt.Printf("The multiple of %v and %v is %v \n", a, b, mult_response)

	div_response := division(a,b)
	fmt.Printf("%v / %v is  %v", a, b, div_response)

	fizz_buzz()
}

func addition(a int, b int)int{
	return a + b
}

func substration(a int, b int)int{
	return a - b
}

func multiplication(a int, b int)int{
	return a * b
}

func division(a int, b int)int{
	return a % b
}

// Fizz buzz

func fizz_buzz(){
	for x:= 1; x <= 100; x++ {
		if x % 3 == 0{
			fmt.Println("Fizz")
		} else if x % 5 == 0{
			fmt.Print("Buzz")
		} else if x % 2 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if x%3 !=0 && x%5 !=0{
				fmt.Println(x)
			}			
		}
		fmt.Printf("\n")
	}
}






