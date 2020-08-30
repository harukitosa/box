package mat

import "fmt"

// Add is add two matrix
func Add(a *Mat, b *Mat) (Mat, error) {
	var c Mat
	if a.row != b.row {
		return c, fmt.Errorf("Argument row %d, %d is not equal", a.row, b.row)
	}
	if a.column != b.column {
		return c, fmt.Errorf("Argument column %d, %d is not equal", a.column, b.column)
	}
	num := a.row * a.column
	c.row = a.row
	c.column = b.column
	for i := 0; i < num; i++ {
		c.value = append(c.value, a.value[i]+b.value[i])
	}
	return c, nil
}
