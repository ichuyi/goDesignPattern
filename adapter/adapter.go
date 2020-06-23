package adapter

import "fmt"

type Adapter interface {
	Do()
}
type AdapterE struct {
	Name string
}

func (a *AdapterE)ToString()  {
	fmt.Println("name: ",a.Name)
}
type RealAdapter struct{
	a *AdapterE
}

func NewRealAdapter(a *AdapterE) *RealAdapter {
	return &RealAdapter{
		a: a,
	}
}
func (a *RealAdapter) Do()  {
	fmt.Println("start visit AdapterE")
	a.a.ToString()
}
