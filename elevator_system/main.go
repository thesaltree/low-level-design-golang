package main

import "sync"

func main() {
	building := NewBuilding()
	manager := NewElevatorManager(building)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		elevator := building.Floors[1].HallPanels[1].RequestElevator(manager, Up)
		elevator.AddDestination(6)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		elevator := building.Floors[8].HallPanels[2].RequestElevator(manager, Down)
		elevator.AddDestination(7)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		thirdElevator := building.Floors[3].HallPanels[0].RequestElevator(manager, Up)
		thirdElevator.AddDestination(12)
	}()

	wg.Wait()

	go manager.OperateAllElevators()

	select {}
}
