package mat

import "fmt"

// Mat is matrix
// Row: 行
// Column: 列
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

// Row return Matrix row
// 行の値を返す
func (m *Mat) Row() int {
	return m.row
}

// Column return Matrix column
// 列の値を返す
func (m *Mat) Column() int {
	return m.column
}

// Value return Matrix data
// 行列の要素を配列にして返す
func (m *Mat) Value() []float64 {
	return m.value
}

// EqualValue returns true if the elements of the matrix
// have equal values and false if they are different.
// It's only a comparison of elements, so it doesn't compare
// whether the objects are identical to each other.
// 行列の要素の値が等しければtrue、異なっていればfalseを返す
// あくまで要素の比較なのでオブジェクト同士が同一のものかどうかは比較しない
func (m *Mat) EqualValue(m2 *Mat) bool {
	if m.Row() != m2.Row() {
		return false
	}
	if m.Column() != m2.Column() {
		return false
	}
	mValue := m.Value()
	m2Value := m2.Value()
	for i := 0; i < len(mValue); i++ {
		if mValue[i] != m2Value[i] {
			return false
		}
	}
	return true
}
