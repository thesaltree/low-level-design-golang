package main

type VendingMachine struct {
	inventory             *Inventory
	currentState          State
	moneyInsertedState    *MoneyInsertedState
	productSelectedState  *ProductSelectedState
	productDispensedState *ProductDispensedState
	selectedProduct       *Product
	insertedMoney         float64
	returnedChange        float64
}

func NewVendingMachine() *VendingMachine {
	vendingMachine := &VendingMachine{inventory: NewInventory()}

	vendingMachine.moneyInsertedState = &MoneyInsertedState{vendingMachine}
	vendingMachine.productSelectedState = &ProductSelectedState{vendingMachine}
	vendingMachine.productDispensedState = &ProductDispensedState{vendingMachine}
	vendingMachine.currentState = &MoneyInsertedState{vendingMachine: vendingMachine}
	return vendingMachine
}

func (v *VendingMachine) InsertMoney(amount float64) {
	v.currentState.InsertMoney(amount)
}

func (v *VendingMachine) SelectProduct(product *Product) {
	v.currentState.SelectProduct(product)
}

func (v *VendingMachine) ReturnChange() {
	v.currentState.ReturnChange()
}

func (v *VendingMachine) DispenseProduct() {
	v.currentState.DispenseProduct()
}

func (v *VendingMachine) UpdateState(newState State) {
	v.currentState = newState
}
