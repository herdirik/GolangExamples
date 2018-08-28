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
	p := Point{3, 4}
	m := MyInt(3)
	
	fmt.Printf("Point: %d\n", addMagicNumber(p))
	fmt.Printf("MyInt: %d\n", addMagicNumber(m))

	author1 := author{
		"Naveen",
        "Ramanathan",
        "Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,

	}
	post1.details()

	//pointer example 
	r1:=rectangle{
		4,
		5,
	}
	fmt.Println("Rectangle is :",r1)
	fmt.Println("Rectangle area is:", r1.Area_by_value())
	fmt.Println("Rectangle area is:", r1.Area_by_reference())
	fmt.Println("Rectangle area is:", (&r1).Area_by_value())
	fmt.Println("Rectangle area is:", (&r1).Area_by_reference())
	//interface example 
	r := rect{
		3, 
		4,
	}
    c := circle{
		5,
	}
	measure(r)
	measure(c)
	
	// error handling example 
	for _,i :=range []int{7,42}{
		if r ,e :=f1(i);e!=nil{
			fmt.Println("f1 failed :", e)
		}else {
			fmt.Println("f1 worked :",r)
		}

	}
	for _,i :=range []int{7,42}{
		if r,e :=f2(i);e!=nil{
			fmt.Println("f2 failed :", e)
		} else {
            fmt.Println("f2 workd :", r)
		}

	}
	_,e :=f2(42)
	if ae , ok := e.(*argError); ok{
		fmt.Println(ae.arg)
        fmt.Println(ae.prob)
	}
	



}