package main

import (
	"designPattern/factory"
)

func main()  {
	c:=factory.ClothesFactory{}
	b:=c.Create()
	b.ToString()
}
