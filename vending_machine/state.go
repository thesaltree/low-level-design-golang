package main

import "fmt"

type State interface {
	InsertMoney(amount float64)
	SelectProduct(product *Product)
	ReturnChange()
	DispenseProduct()
}

type MoneyInsertedState struct {
	vendingMachine *VendingMachine
}

func (s *MoneyInsertedState) InsertMoney(amount float64) {
	s.vendingMachine.insertedMoney += amount
	fmt.Printf("Inserted amount: $%.2f\n", amount)
	s.vendingMachine.UpdateState(s.vendingMachine.productSelectedState)
}

func (s *MoneyInsertedState) SelectProduct(product *Product) {
	fmt.Println("Insert money to proceed")
}

func (s *MoneyInsertedState) ReturnChange() {
	if s.vendingMachine.insertedMoney > 0 {
		change := s.vendingMachine.insertedMoney
		s.vendingMachine.insertedMoney = 0
		fmt.Printf("Returning change: $%.2f\n", change)
	} else {
		fmt.Println("No change to return.")
	}
}

func (s *MoneyInsertedState) DispenseProduct() {
	fmt.Println("Insert money to purchase product")
}

type ProductSelectedState struct {
	vendingMachine *VendingMachine
}

func (s *ProductSelectedState) InsertMoney(amount float64) {
	fmt.Println("Product already selected, please wait.")
}

func (s *ProductSelectedState) SelectProduct(product *Product) {
	if !s.vendingMachine.inventory.IsProductAvailable(product.ID) {
		fmt.Println("Product not available.")
		s.vendingMachine.UpdateState(s.vendingMachine.moneyInsertedState)
		return
	}

	if product.Price > s.vendingMachine.insertedMoney {
		fmt.Printf("Insufficient funds. Inserted: $%.2f, Price: $%.2f\n", s.vendingMachine.insertedMoney, product.Price)
		s.vendingMachine.UpdateState(s.vendingMachine.moneyInsertedState)
	} else {
		s.vendingMachine.selectedProduct = product
		fmt.Printf("Product %s selected.\n", product.Name)
		s.vendingMachine.UpdateState(s.vendingMachine.productDispensedState)
	}
}

func (s *ProductSelectedState) ReturnChange() {
	s.vendingMachine.moneyInsertedState.ReturnChange()
}

func (s *ProductSelectedState) DispenseProduct() {
	fmt.Println("Cannot dispense. Select product first.")
}

type ProductDispensedState struct {
	vendingMachine *VendingMachine
}

func (s *ProductDispensedState) InsertMoney(amount float64) {
	fmt.Println("Please wait. Product is being dispensed.")
}

func (s *ProductDispensedState) SelectProduct(product *Product) {
	fmt.Println("Product already dispensed.")
}

func (s *ProductDispensedState) ReturnChange() {
	s.vendingMachine.moneyInsertedState.ReturnChange()
}

func (s *ProductDispensedState) DispenseProduct() {
	product, err := s.vendingMachine.inventory.TransactProduct(s.vendingMachine.selectedProduct.ID)
	if err != nil {
		fmt.Println("Error dispensing product:", err)
		s.vendingMachine.UpdateState(s.vendingMachine.moneyInsertedState)
		return
	}

	change := s.vendingMachine.insertedMoney - product.Price
	s.vendingMachine.insertedMoney = 0
	if change > 0 {
		fmt.Printf("Returning change: $%.2f\n", change)
	}
	fmt.Printf("Dispensed product: %s\n", product.Name)
	s.vendingMachine.UpdateState(s.vendingMachine.moneyInsertedState)
}
