package engine

type SlotMachine interface {
	SpinOnce() Field
	LinesCount() int
}

type Field interface {
	TotalPrize(bet int64, linesCount int) (int64, error)
	BonusGames() int
	GetStops() []int
	Prize(lineNum int) (int64, error)
}
