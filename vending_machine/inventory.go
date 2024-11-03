package main

import "fmt"

type Inventory struct {
	Products map[int]*Product
}

func NewInventory() *Inventory {
	return &Inventory{Products: make(map[int]*Product)}
}

func (i *Inventory) AddProduct(product *Product) {
	i.Products[product.ID] = product
	fmt.Printf("Added product with ID %d\n", product.ID)
}

func (i *Inventory) DeleteProduct(id int) {
	delete(i.Products, id)
	fmt.Printf("Deleted product with ID %d\n", id)
}

func (i *Inventory) TransactProduct(id int) (*Product, error) {
	product, ok := i.Products[id]
	if !ok {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}
	if product.Quantity > 0 {
		product.Quantity--
		if product.Quantity == 0 {
			delete(i.Products, id)
		}
		fmt.Printf("Transacted product with ID %d, remaining quantity: %d\n", id, product.Quantity)
		return product, nil
	}
	return nil, fmt.Errorf("product with ID %d is out of stock", id)
}

func (i *Inventory) IsProductAvailable(id int) bool {
	product, ok := i.Products[id]
	return ok && product.Quantity > 0
}
