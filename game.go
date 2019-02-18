package main

import (
	"slot_machine/engine"
	"slot_machine/engine/atkins"
)

type slotMachine struct {
	machine engine.SlotMachine
}

type playResult struct {
	Type  string
	Total int64
	Stops []int
}

func newAtkinsGame() slotMachine {
	return slotMachine{
		machine: atkins.NewMachine(nil),
	}
}

func (g slotMachine) play(bet int64) ([]playResult, error) {
	result := make([]playResult, 0, 1)

	games := 1
	bonusGame := false

	for ; games > 0; games-- {
		f := g.machine.SpinOnce()
		prize, err := f.TotalPrize(bet, 1)
		if err != nil {
			return nil, err
		}

		gameType := "main"
		if bonusGame {
			gameType = "free"
		}
		result = append(result, playResult{
			Type:  gameType,
			Total: prize,
			Stops: f.GetStops(),
		})

		games += f.BonusGames()
		bonusGame = true
	}

	return result, nil
}
