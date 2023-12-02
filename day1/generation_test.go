package day1

import "testing"

func TestBitWidth(t *testing.T) {
	cases := []struct {
		n     int
		width int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 3},
		{7, 3},
	}
	for _, c := range cases {
		w := BitWidth(c.n)
		if w != c.width {
			t.Errorf("BitWidth(%d) == %d, expected %d", c.n, w, c.width)
		}
	}
}
