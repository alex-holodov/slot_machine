package atkins

import (
	"testing"
)

func TestScaleCount(t *testing.T) {
	tests := []struct {
		values [RelesCount]int
		scales int
	}{
		{[RelesCount]int{0, 0, 0, 0, 0}, 2},
		{[RelesCount]int{19, 13, 25, 2, 18}, 1},
		{[RelesCount]int{0, 21, 0, 0, 0}, 3},
		{[RelesCount]int{8, 21, 1, 0, 0}, 2},
		{[RelesCount]int{6, 6, 6, 6, 6}, 0},
	}

	for _, test := range tests {
		r := &testRandom{}
		r.setTestValues(test.values)

		e := NewMachine(r)
		field := e.SpinOnce()

		if field.(*gameField).scales != test.scales {
			t.Errorf("bad scale value for %#v - %d, wanted %d", test.values, field.(*gameField).scales, test.scales)
		}
	}
}

func TestEngine(t *testing.T) {
	tests := []struct {
		values [RelesCount]int
		result int64
		bonus  int
	}{
		{[RelesCount]int{0, 0, 0, 0, 0}, 0, 0},
		{[RelesCount]int{19, 13, 25, 2, 18}, 5000, 0},
		{[RelesCount]int{0, 21, 0, 0, 0}, 0, 10},
	}

	for _, test := range tests {
		r := &testRandom{}
		r.setTestValues(test.values)

		e := NewMachine(r)
		field := e.SpinOnce()

		result, err := field.Prize(0)
		if err != nil {
			t.Fatalf("can't get prize: %s", err)
		}
		if result != test.result {
			t.Errorf("bad result value for %#v - %d, wanted %d", test.values, result, test.result)
		}

		bonus := field.BonusGames()
		if bonus != test.bonus {
			t.Errorf("bad bonus value for %#v - %d, wanted %d", test.values, bonus, test.bonus)
		}
	}
}
