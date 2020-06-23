package abstractFactory

import "fmt"

type AbstractFactory interface {
	CreateShop() Shop
	CreateFood() Food
}

type Shop interface {
	ToString()
}
type Clothes struct {
	Price int
	Name string
}
func (c *Clothes)ToString()  {
	fmt.Printf("name is %s, price is %d\n",c.Name,c.Price)
}
type Food interface {
	Eat()
}
type Meat struct {
	Weight int
}

func (m *Meat)Eat()  {
	fmt.Printf("meat weight %d\n",m.Weight)
}

type RealizeFactory struct{

}

func (c *RealizeFactory)CreateShop()Shop  {
	return &Clothes{
		Price: 100,
		Name:  "",
	}
}
func (c *RealizeFactory)CreateFood()Food  {
	return &Meat{
		Weight: 10,
	}
}