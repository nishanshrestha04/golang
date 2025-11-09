package main

import "fmt"
func main(){
	calculator := func(a int, b int, opt string){
		switch opt{
		case "*":
			fmt.Println("The product is ", a * b)
		case "/":
			fmt.Println("The division is ", a / b)
		case "+":
			fmt.Println("The addition is ", a + b)
		case "-":
			fmt.Println("The subtrction is ", a - b)
		}
	}

	calculator(1, 556, "+")
}