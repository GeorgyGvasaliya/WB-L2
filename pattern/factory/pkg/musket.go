package pkg

type musket struct {
	gun
}

func newMusket() IGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
