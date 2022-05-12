package pkg

type builder struct {
	windowType string
	doorType   string
	floor      int
}

func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *builder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *builder) setNumFloor() {
	b.floor = 2
}

func (b *builder) getHouse() house {
	return house{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
