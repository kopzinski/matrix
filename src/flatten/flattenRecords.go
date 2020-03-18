package flatten

func FlattenRecords(records [][]int) []int {

	flatten := []int{}

	for _, row := range records {
		flatten = append(flatten, row...)
	}

	return flatten
}
