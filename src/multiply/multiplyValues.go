package multiply

import (
	"calc"
)

func MultiplyValues(records [][]int) int {
	return calc.BulkCalculate(records, calc.Multiply)
}
