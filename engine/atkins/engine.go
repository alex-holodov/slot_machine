package atkins

import (
	"slot_machine/engine"
)

type machine struct {
	random Random
}

func NewMachine(random Random) *machine {
	if random == nil {
		random = DefaultRandom{}
	}

	return &machine{
		random: random,
	}
}

func (m *machine) SpinOnce() engine.Field {
	stops := m.random.GetRandom()
	field := m.prepareField(stops)
	scales := getScaleCount(field)

	return &gameField{
		symbols: field,
		stops:   stops,
		scales:  scales,
	}
}

func (m *machine) LinesCount() int {
	return PayLinesCount
}

func (m *machine) prepareField(stops [RelesCount]int) [ViewableLines][RelesCount]Symbol {
	var result [ViewableLines][RelesCount]Symbol

	for i, pos := range stops {
		up, down := m.getNeighbours(pos)

		result[0][i] = getSymbol(up, i)
		result[1][i] = getSymbol(pos, i)
		result[2][i] = getSymbol(down, i)
	}

	return result
}

func getScaleCount(field [ViewableLines][RelesCount]Symbol) int {
	scales := 0
	for _, l := range field {
		for _, s := range l {
			if s == Scale {
				scales++
			}
		}
	}

	return scales
}

func (m *machine) getNeighbours(pos int) (up int, down int) {
	up = pos - 1
	if up < 0 {
		up = TotalLines - 1
	}

	down = pos + 1
	if down == TotalLines {
		down = 0
	}

	return
}

func getSymbol(lineNum int, releNum int) Symbol {
	return stripsTable[lineNum][releNum]
}
