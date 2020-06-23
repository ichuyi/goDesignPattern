package singleton

import (
	"fmt"
	"sync"
)

type WangSonghao struct {
	Name string
	Age int
}
var p *WangSonghao
var once *sync.Once
var Count int

func init()  {
	once=new(sync.Once)
	fmt.Println(once)
}
func getPeople()  {
	Count++
	p=&WangSonghao{
		Name: "hhh",
		Age: 0,
	}
}
func GetPeople()*WangSonghao  {
	once.Do(func() {
		getPeople()
	})
	fmt.Println(once)
	return p
}
