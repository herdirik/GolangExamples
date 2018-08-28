package main 

type rectangle struct {
	length , width int
}
func (r rectangle) Area_by_value() int{
	return r.length * r.width
}
func (r *rectangle) Area_by_reference() int{
	return r.length * r.width
}