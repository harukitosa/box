package mat

import (
	"testing"
)

func TestNewMatValid(t *testing.T) {
	testCases := []struct {
		width  int
		height int
		value  []float64
	}{
		{3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{2, 3, []float64{1, 1, 1, 1, 2, 2}},
	}
	for _, test := range testCases {
		m, err := NewMat(test.value, test.width, test.height)
		if err != nil {
			t.Errorf("error msg:%v", err)
		}
		for i := 0; i < test.width*test.height; i++ {
			if m.value[i] != test.value[i] {
				t.Errorf("error msg")
			}
		}
	}
}

func TestNewMatError(t *testing.T) {
	testCases := []struct {
		width  int
		height int
		value  []float64
	}{
		{3, 4, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{10, 3, []float64{1, 1, 1, 1, 2, 2}},
	}
	for _, test := range testCases {
		_, err := NewMat(test.value, test.width, test.height)
		if err == nil {
			t.Errorf("test fails width ans height is not valid value")
		}
	}
}
