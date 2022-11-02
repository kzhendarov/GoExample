package constants

// Пример реализации собственного типа данных с реализацией перечисления средсвом набора констант.
type Direction byte

const (
	// iota - автоинкримент целочисленного типа данных в рамках одного блока определения значений
	North Direction = iota // тип Direction(byte) = 0
	East                   // тип Direction(byte) = 1
	South                  // тип Direction(byte) = 2
	West                   // тип Direction(byte) = 3
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// : Выбор направления от значения экземпляра типа Direction
func (d Direction) DirectionSelect() (s string) {
	switch d {
	case North:
		s = "Move UP"
	case East:
		s = "Move RIGHT"
	case South:
		s = "Move DOWN"
	case West:
		s = "Move LEFT"
	}
	return s
}
