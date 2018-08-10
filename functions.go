package main

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