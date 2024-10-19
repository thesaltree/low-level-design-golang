package main

import (
	"fmt"
	"lld_go_parking_lot/vehicles"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	parkingLot := GetParkingLotInstance()

	parkingLot.Name = "Central Parking Lot"

	parkingLot.AddFloor(0)
	parkingLot.AddFloor(1)

	parkingLot.DisplayAvailability()

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go parkCar(i, parkingLot)
	}

	wg.Wait()

	parkingLot.DisplayAvailability()

	ticket, _ := parkingLot.ParkVehicle(vehicles.NewTruck("truck-1"))

	time.Sleep(10 * time.Second)
	err := parkingLot.UnparkVehicle(ticket)
	if err != nil {
		return
	}

	formattedCharge := fmt.Sprintf("%.2f", ticket.CalculateTotalCharge())

	fmt.Printf("bill for %s = %s\n", ticket.Vehicle.GetLicenceNumber(), formattedCharge)

}

func parkCar(ind int, parkingLot *ParkingLot) {
	defer wg.Done()

	car := vehicles.NewCar(fmt.Sprintf("car-%d", ind))

	ticket, err := parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Printf("Failed to park %s: %v\n", car.LicenceNumber, err)
		return
	}

	fmt.Printf("%s parked successfully. Ticket: %s\n", car.LicenceNumber, ticket.EntryTime)
}
