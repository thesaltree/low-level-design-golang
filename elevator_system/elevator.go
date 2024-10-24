package main

import "fmt"

type Elevator struct {
	ID               int
	Capacity         int
	CurrentFloor     int
	CurrentDirection Directions
	CurrentLoad      int
	ElevatorPanel    *ElevatorPanel
	Destinations     []int
}

func NewElevator(id int) *Elevator {
	return &Elevator{ID: id, Capacity: 10, CurrentFloor: 1, CurrentDirection: Still, CurrentLoad: 0, ElevatorPanel: NewElevatorPanel(id)}
}

func (e *Elevator) AddDestination(destinationFloor int) {
	e.ElevatorPanel.AddDestinationFloor(destinationFloor)
	e.Destinations = append(e.Destinations, destinationFloor)
	fmt.Printf("Elevator %d received destination floor %d\n", e.ID, destinationFloor)
}

func (e *Elevator) RemoveDestination(destinationFloor int) {
	for i, floor := range e.Destinations {
		if floor == destinationFloor {
			e.Destinations = append(e.Destinations[:i], e.Destinations[i+1:]...)
			e.ElevatorPanel.RemoveDestinationFloor(destinationFloor)
			break
		}
	}
}

func (e *Elevator) UpdateCurrentFloor(newFloor int) {
	e.CurrentFloor = newFloor
}

func (e *Elevator) UpdateCurrentLoad(newLoad int) {
	e.CurrentLoad = newLoad
}

func (e *Elevator) UpdateCurrentDirection(newDirection Directions) {
	e.CurrentDirection = newDirection
}
