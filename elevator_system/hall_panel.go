package main

import "fmt"

type Directions string

const (
	Up    Directions = "Up"
	Down  Directions = "Down"
	Still Directions = "Still"
)

type HallPanel struct {
	PanelID              int
	DirectionInstruction Directions
	SourceFloor          int
}

func NewHallPanel(panelID int, sourceFloor int) *HallPanel {
	return &HallPanel{PanelID: panelID, SourceFloor: sourceFloor, DirectionInstruction: Still}
}

func (h *HallPanel) SetDirectionInstructions(directionInstruction Directions) {
	h.DirectionInstruction = directionInstruction
}

func (h *HallPanel) RequestElevator(manager *ElevatorManager, direction Directions) (elevator *Elevator) {
	fmt.Printf("Panel %d requesting elevator with direction %s from floor %d\n", h.PanelID, direction, h.SourceFloor)
	return manager.AssignElevator(h.SourceFloor, direction)
}
