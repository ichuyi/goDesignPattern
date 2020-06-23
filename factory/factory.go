package factory

import "fmt"

type Factory interface {
	Create() Shop
}

type Shop interface {
	ToString()
}
type Clothes struct {
	Price int
	Name string
}

func (c *Clothes) ToString()  {
	fmt.Printf("name is %s, price is %d\n",c.Name,c.Price)
}
type ClothesFactory struct{

}

func (c *ClothesFactory)Create()Shop  {
	return &Clothes{
		Price: 100,
		Name:  "",
	}
}