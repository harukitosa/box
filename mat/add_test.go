package mat

import (
	"testing"
)

func TestAddValid(t *testing.T) {
	testCases := []struct {
		width  int
		height int
		value  []float64
	}{
		{3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{3, 3, []float64{1, 1, 1, 2, 2, 2, 3, 3, 3}},
	}

	answer := []float64{2, 2, 2, 3, 3, 3, 4, 4, 4}
	ans, _ := NewMat(answer, 3, 3)
	a, _ := NewMat(testCases[0].value, testCases[0].width, testCases[0].height)
	b, _ := NewMat(testCases[1].value, testCases[1].width, testCases[1].height)
	result, err := Add(a, b)
	if err != nil {
		t.Errorf("test fails")
	}
	if !ans.EqualValue(&result) {
		t.Errorf("Error at Add method")
	}

}
