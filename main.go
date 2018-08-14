package main

import (
	"fmt"
	"encoding/json"


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
	//for döngüsü.
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
	//Copy (kopyalama) örneği.
	kaynak:=[]int{1,2,3,4,5}
	hedef:=make([]int, 2)
	fmt.Println(hedef, kaynak)
	sayi:=copy(hedef,kaynak)
	fmt.Println("Kopyalanan eleman sayısı:" ,sayi)
	fmt.Println(hedef, kaynak)
	//Map(Harita) örneği.
	var il_nufus map[string]int
    il_nufus = make(map[string]int)
	il_nufus["İstanbul"]=1500000
	il_nufus["izmir"]=35.000
	il_nufus["Manisa"]=20.000
	fmt.Println(il_nufus["Manisa"])
	fmt.Println(il_nufus)

	sozluk := make(map[string]string)
	sozluk["IP"]="Internet Protocol"
	sozluk["LAMP"]="Linux Apache MySQL PHP"
	sozluk["GNU"] ="GNU is not Unix"
	sozluk["WWW"] ="World Wide Web"
	sozluk["IMAP"] ="Internet Message Access Protocol"
	tumelemanlar,_:=json.MarshalIndent(sozluk,""," ")
	fmt.Println(string(tumelemanlar))

	fmt.Println(faktoriyel(5))



}