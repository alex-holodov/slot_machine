package atkins

import "fmt"

type gameField struct {
	symbols [LinesAmount][RelesCount]Symbol
	stops   [RelesCount]int
	scales  int
}

func (g *gameField) GetStops() []int {
	stops := make([]int, len(g.stops))
	for i, s := range g.stops {
		stops[i] = s + 1
	}
	return stops
}

func (g *gameField) TotalPrize(bet int64, linesCount int) (int64, error) {
	if linesCount > PayLines {
		return 0, fmt.Errorf("bad line count")
	}

	var total int64
	for i := 0; i < linesCount; i++ {
		prize, err := g.Prize(i)
		if err != nil {
			return 0, err
		}
		total += prize
	}

	total += scalePayTable[g.scales]

	return total * bet, nil
}

func (g *gameField) Prize(lineNum int) (int64, error) {
	if lineNum >= PayLines {
		return 0, fmt.Errorf("bad line number")
	}

	l := g.getLine(lineNum)
	return calculatePrize(l), nil
}

func (g *gameField) BonusGames() int {
	if g.scales >= MinScalesToBonus {
		return BonusGames
	}

	return 0
}

func (g *gameField) getLine(lineNo int) [RelesCount]Symbol {
	line := lineTable[lineNo]

	var result [RelesCount]Symbol
	for i, l := range line {
		result[i] = g.symbols[l][i]
	}

	return result
}

func calculatePrize(line [RelesCount]Symbol) int64 {
	var inRow int
	i := 0

	// skip first Atkins symbols
	for ; i < RelesCount; i++ {
		if line[i] != Atkins {
			break
		}
		inRow++
	}
	atkinsPrize := getPrize(Atkins, inRow)
	if i == RelesCount {
		return atkinsPrize
	}

	sym := line[i]
	for ; i < RelesCount; i++ {
		if sym != line[i] && line[i] != Atkins {
			break
		}
		inRow++
	}
	prize := getPrize(sym, inRow)
	if atkinsPrize > prize {
		return atkinsPrize
	}

	return prize
}

func getPrize(sym Symbol, inRow int) int64 {
	if inRow == 0 {
		return 0
	}

	return payTable[sym][RelesCount-inRow]
}
