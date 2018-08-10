package main

import (
	"fmt"

)
func main(){
	
	var eposta , password , age bool
	var number1, number2 ,sum int
	var portion float32
	var name, surname string
	const country ="TR"
	
	eposta=true
	password=true
	age=false
	name= "Hatice"
	surname = "\nErdirik"
	number1 = 3
	number2 = number1+5
	sum= number1 + number2
	portion = 3.0 / 8.0

	fmt.Println(sum)
	fmt.Println(portion)
	fmt.Println(name+surname)
	fmt.Println(eposta && age)
	fmt.Println(password || age)
	fmt.Println(!eposta)
	fmt.Println(country)
	fmt.Println(add(3,5))
	fmt.Println(split(9))
	a,b :=swap("hello","world")
	fmt.Println(a,b)
}