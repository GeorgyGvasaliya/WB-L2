package pkg

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "wooden" {
		return &builder{}
	}

	return nil
}
