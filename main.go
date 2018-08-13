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
	var poww= []int{1,2,4,8,16,32,64,128}
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
    s:=0
	for i:=0 ;i<10;i++{
		s+=i
	}
	fmt.Println(s)
	fmt.Println(sm(1))
	//while deyimin karşılığı go da for dur.
	number:=1
	for number < 1000{
		number+=number
	}
	fmt.Println(number)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	
	for i, v := range poww {
			
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// slices example 
	slice :=make([] string ,3)
	fmt.Println("emp:",slice)
	slice[0]="a"
	slice[1]="b"
	slice[2]="c"
	fmt.Println("set:",slice)
	fmt.Println("len:", len(slice))


}