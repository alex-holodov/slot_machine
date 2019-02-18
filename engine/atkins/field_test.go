package atkins

import "testing"

func TestCalculatePrize(t *testing.T) {
	tests := []struct {
		values [RelesCount]Symbol
		result int64
	}{
		{[RelesCount]Symbol{Atkins, Atkins, Atkins, Atkins, Atkins}, 5000},
		{[RelesCount]Symbol{Ham, Bacon, Buffalo, Bacon, Atkins}, 0},
		{[RelesCount]Symbol{Ham, Bacon, Buffalo, Bacon, Atkins}, 0},
		{[RelesCount]Symbol{Atkins, Mayonnaise, Ham, Ham, Ham}, 0},
		{[RelesCount]Symbol{Atkins, Ham, Ham, Ham, Ham}, 500},
		{[RelesCount]Symbol{Atkins, Atkins, Ham, Ham, Ham}, 500},
		{[RelesCount]Symbol{Atkins, Ham, Atkins, Ham, Atkins}, 500},
		{[RelesCount]Symbol{Atkins, Ham, Atkins, Atkins, Atkins}, 500},
		{[RelesCount]Symbol{Atkins, Atkins, Ham, Atkins, Atkins}, 500},
		{[RelesCount]Symbol{Atkins, Atkins, Ham, Atkins, Atkins}, 500},
		{[RelesCount]Symbol{Atkins, Atkins, Atkins, Atkins, Mayonnaise}, 500},
		{[RelesCount]Symbol{Atkins, Atkins, Atkins, Mayonnaise, Mayonnaise}, 50},
		{[RelesCount]Symbol{Mayonnaise, Mayonnaise, Mayonnaise, Mayonnaise, Mayonnaise}, 50},
	}

	for _, te := range tests {
		result := calculatePrize(te.values)
		if result != te.result {
			t.Errorf("bad result %d, wanted %d for %#v", result, te.result, te.values)
		}
	}
}
