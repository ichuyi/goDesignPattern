package proxy

import "fmt"

type Object interface {
	Do()
}
type RealObject struct{

}

func (r *RealObject)Do()  {
	fmt.Println("real object do")
}
type Proxy struct {
	r RealObject
}

func (r *Proxy)Do()  {
	r.preDo()
	r.r.Do()
	r.afterDo()
}
func (r *Proxy)preDo()  {
	fmt.Println("proxy pre do")
}
func (r *Proxy) afterDo()  {
	fmt.Println("proxy after do")
}
