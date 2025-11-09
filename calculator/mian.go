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
	var num1 int
	var num2 int
	var opt string
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)	
	fmt.Print("Enter the operation you want to perform: ")
	fmt.Scan(&opt)

	calculator(num1, num2, opt)
}