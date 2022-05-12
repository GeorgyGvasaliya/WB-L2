package pkg

type director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
