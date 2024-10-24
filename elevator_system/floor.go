package main

type Floor struct {
	Number     int
	HallPanels []*HallPanel
}

func NewFloor(number int) *Floor {
	floor := &Floor{Number: number, HallPanels: make([]*HallPanel, 0)}

	for i := 1; i <= 3; i++ {
		hallPanel := NewHallPanel(i, number)
		floor.HallPanels = append(floor.HallPanels, hallPanel)
	}

	return floor
}
