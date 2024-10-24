package main

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
