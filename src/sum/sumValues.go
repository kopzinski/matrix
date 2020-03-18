package sum

import (
	"calc"
)

func SumValues(records [][]int) int {
	return calc.BulkCalculate(records, calc.Sum)
}
