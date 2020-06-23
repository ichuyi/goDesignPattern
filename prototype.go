package main

import (
	"designPattern/prototype"
	"fmt"
)

func main()  {
	p:=prototype.WangSonghao.Clone()
	switch p.(type) {
	case *prototype.RealizationPrototype:
		fmt.Println(p.(*prototype.RealizationPrototype))
	default:
		fmt.Println("default")
	}
}
