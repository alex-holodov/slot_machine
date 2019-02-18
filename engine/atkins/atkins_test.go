package atkins

import "testing"

func TestAtkins(t *testing.T) {
	tests := []struct {
		values     [RelesCount]int
		bet        int64
		linesCount int
		result     int64
	}{
		{[RelesCount]int{26, 13, 2, 30, 26}, 5, 20, 75},
		{[RelesCount]int{15, 26, 8, 18, 31}, 10, 1, 20},
	}

	for _, test := range tests {
		r := &testRandom{}
		r.setTestValues(test.values)

		e := NewMachine(r)
		field := e.SpinOnce()
		prize, err := field.TotalPrize(test.bet, test.linesCount)
		if err != nil {
			t.Fatalf("get prize error: %s", err)
		}

		if prize != test.result {
			t.Errorf("bad total prize value %d, watned %d, (%+v)", prize, test.result, test.values)
		}

	}

}
