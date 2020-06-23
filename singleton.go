package main

import (
	"designPattern/singleton"
	"fmt"
)

func main()  {
	p:=singleton.GetPeople()
	fmt.Println(p)
	fmt.Println(singleton.Count)
	q:=singleton.GetPeople()
	fmt.Println(q)
	fmt.Println(singleton.Count)
}
