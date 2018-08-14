package main
import(
	"math"
)

func add(x int , y int) int{
    return x + y
}

func split(sum int)(x,y int){
	x=sum * 4/9
	y=sum -x
	return 
}

func swap(x,y string)(string ,string){
	return y,x
}
func sm(sum int) int{
	for ; sum <1000;{
		sum+=sum

	}
	return sum
}
//if içerisinde olan kısa atama.
func pow(x, n, lim float64) float64{
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
//Recursive fonksiyonlar.
func faktoriyel(sayi int) int{
	if sayi ==0{
		return 1
	}
	return sayi * faktoriyel(sayi -1)
}
