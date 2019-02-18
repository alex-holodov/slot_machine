package atkins

import "math/rand"

type Random interface {
	GetRandom() [RelesCount]int
}

type DefaultRandom struct{}

func (r DefaultRandom) GetRandom() [RelesCount]int {
	var result [RelesCount]int
	for i := 0; i < RelesCount; i++ {
		result[i] = int(rand.Int31n(TotalLines))
	}

	return result
}

type testRandom struct {
	values [RelesCount]int
}

func (t *testRandom) GetRandom() [RelesCount]int {
	return t.values
}

func (t *testRandom) setTestValues(v [RelesCount]int) {
	t.values = v
}
