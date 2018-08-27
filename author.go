package main
import (
	"fmt"
)
type author struct{
	firtname string
	lastname string
	bio string

}
func (a author) fullname() string{
	return fmt.Sprintf("%s %s",a.firtname ,a.lastname)
}
type post struct{
	title string
	content string
	author
}
func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.fullname())
    fmt.Println("Bio: ", p.bio)
}