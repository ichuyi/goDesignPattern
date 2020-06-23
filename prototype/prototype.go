package prototype

import "sync"

type Prototype interface {
	Clone() Prototype
}

type RealizationPrototype struct {
	Name string
	Age int
}

func (r *RealizationPrototype)Clone() Prototype  {
	p:= RealizationPrototype{
		Name: r.Name,
		Age: r.Age,
	}
	return &p
}
var WangSonghao *RealizationPrototype
var once *sync.Once
func init()  {
	once=new(sync.Once)
	once.Do(func() {
		WangSonghao=new(RealizationPrototype)
	})
}