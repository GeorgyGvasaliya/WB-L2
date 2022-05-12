package pkg

type department interface {
	execute(patient *Patient)
	SetNext(department)
}
