package product

type Basket struct {
	Name string
	Products []Product
}

func (b *Basket)  AddProduct(p Product)  {
    b.Products =  append(b.Products, p)


/*func (b *Basket) PrintProducts() {
	fmt.Println(b.Products)

	}
*/
}


