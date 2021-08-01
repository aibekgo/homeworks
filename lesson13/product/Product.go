package product

import (
	"fmt"
)

type Product struct {
	Name string
	Price int
}

func (p *Product) GetFullInfo() string{
	//result := fmt.Sprintf( "Name:%s; Price: %d",p.Price, p.Name )
	result := fmt.Sprintf( "Name:%s", p.Name )
	return result
}



//Product:
/*Name:string
Price:int
GetFullInfo():string*/

