package main

func main() {
	vendingMachine := NewVendingMachine()

	// Add products to the inventory
	coke := NewProduct(1, "Coke", 1.50)
	pepsi := NewProduct(2, "Pepsi", 1.75)

	vendingMachine.inventory.AddProduct(coke)
	vendingMachine.inventory.AddProduct(pepsi)

	// Insert money into the vending machine
	vendingMachine.InsertMoney(2.00)

	// Select a product to purchase
	vendingMachine.SelectProduct(coke)

	// Dispense the selected product
	vendingMachine.DispenseProduct()

	// Return any remaining change
	vendingMachine.ReturnChange()

	// Insert more money to purchase another product with insufficient funds
	vendingMachine.InsertMoney(1.50)

	// Select a second product to purchase
	vendingMachine.SelectProduct(pepsi)

	// Dispense the selected product
	vendingMachine.DispenseProduct()

	// Return any remaining change
	vendingMachine.ReturnChange()

}
