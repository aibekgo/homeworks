package product

import "fmt"

type Basket struct {
	Name string
	Products []Product
}


func (b *Basket)  AddProduct(p Product)  {
    b.Products =  append(b.Products, p) }

func (b *Basket) ListProducts() {
		for i := 0; i < len(b.Products); i++ {
			fmt.Println(b.Products[i].GetFullInfo())
		}
	}


