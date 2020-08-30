package mat

import "fmt"

// Mat is matrix
type Mat struct {
	value  []float64
	column int
	row    int
}

// NewMat is constructor of Mat
func NewMat(value []float64, column int, row int) (*Mat, error) {
	if len(value) != column*row {
		return nil, fmt.Errorf("The values given are not correct for the columns and rows of the matrix")
	}
	if row == 0 {
		return nil, fmt.Errorf("The row is zero")
	}
	if column == 0 {
		return nil, fmt.Errorf("The column is zero")
	}
	m := new(Mat)
	m.value = value
	m.column = column
	m.row = row
	return m, nil
}

// Print is print matrix value
func (m *Mat) Print() {
	for i := 0; i < m.column*m.row; i++ {
		fmt.Printf("%5.8f ", m.value[i])
		if (i+1)%m.row == 0 {
			fmt.Printf("\n")
		}
	}
}
