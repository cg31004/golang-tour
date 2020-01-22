package main

import "fmt"

type Cart struct {
	Name  string
	Price int
}

func (c Cart) GetPrice() {
	fmt.Println("price: ", c.Price)
}

func (c Cart) UpdatePrice(price int) *Cart {
	c.Price = price
	return &c
}

func (c *Cart) UpdatePricePointer(price int) {
	fmt.Println("[pointer] Update Price to", price)
	c.Price = price
}

func main() {
	c := &Cart{"bage", 100}
	c.GetPrice()
	c.UpdatePrice(200)
	fmt.Println(c)
	c.UpdatePricePointer(200)
	fmt.Println(c)

}
