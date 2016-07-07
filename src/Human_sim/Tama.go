package Human_sim

type Character interface {
	Act()
}

type Personality struct {
	intelligence int
	charisma     int
	energy       int
	laziness     int
	Disposition
}
