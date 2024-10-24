package main

type Building struct {
	Floors    []*Floor
	Elevators []*Elevator
}

func NewBuilding() *Building {
	building := &Building{Floors: make([]*Floor, 0)}

	for i := 1; i <= 15; i++ {
		floor := NewFloor(i)
		building.Floors = append(building.Floors, floor)
	}

	for i := 1; i <= 3; i++ {
		elevator := NewElevator(i)
		building.Elevators = append(building.Elevators, elevator)
	}

	return building
}
