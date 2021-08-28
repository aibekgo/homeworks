package products

type Product struct {
	Id       string `json:"id"`
	Name 	 string `json:"name"`
	Price    int    `json:"price"`
}

type ProductStore interface {
	Create(user *Product) (*Product, error)
	List() ([]Product, error)
	Delete(id string) error
}
