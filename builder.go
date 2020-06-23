package main

import (
	builder2 "designPattern/builder"
	"fmt"
)

func main()  {
	builder:=new(builder2.ManBuilder)
	director:=new(builder2.Director)
	director.SetBuilder(builder)
	p:=director.Construct()
	switch  p.(type){
	//case builder2.People:
	//	fmt.Println("this is people")
	case *builder2.Man:
		fmt.Println("this is a man")
		fmt.Println(p)
	default:
		fmt.Println("this is nothing")
	}
}
