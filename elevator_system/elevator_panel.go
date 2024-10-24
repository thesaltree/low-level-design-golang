package main

type ElevatorPanel struct {
	PanelID      int
	FloorButtons [15]bool
}

func NewElevatorPanel(panelID int) *ElevatorPanel {
	return &ElevatorPanel{PanelID: panelID, FloorButtons: [15]bool{}}
}

func (ep *ElevatorPanel) AddDestinationFloor(floor int) {
	ep.FloorButtons[floor] = true
}

func (ep *ElevatorPanel) RemoveDestinationFloor(floor int) {
	ep.FloorButtons[floor] = false
}
