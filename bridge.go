package main

import "designPattern/bridge"

func main()  {
	var b bridge.Bag=&bridge.MoneyBag{}
	var r bridge.Color=&bridge.Red{}
	var s bridge.Size=&bridge.Small{}
	b.SetColor(r)
	b.SetSize(s)
	b.Description()
}
