package main

import "designPattern/abstractFactory"

func main()  {
	f:=abstractFactory.RealizeFactory{}
	f.CreateShop().ToString()
	f.CreateFood().Eat()
}
