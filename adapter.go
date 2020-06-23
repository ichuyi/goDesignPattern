package main

import "designPattern/adapter"

func main()  {
	r:=&adapter.AdapterE{
		Name: "wangsonghao",
	}
	var a adapter.Adapter=adapter.NewRealAdapter(r)
	a.Do()
}
