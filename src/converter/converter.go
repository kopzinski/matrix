package converter

import "strconv"

func ConvertStringMatrixToInt(recordsAsString [][]string) ([][]int, error) {
	length := len(recordsAsString)
	recordsAsInt := make([][]int, length)
	for i := range recordsAsInt {
		recordsAsInt[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			valueAsInt, err := strconv.Atoi(recordsAsString[i][j])
			if err != nil {
				return nil, err
			}
			recordsAsInt[i][j] = valueAsInt
		}
	}

	return recordsAsInt, nil
}

func ConvertIntArrayToString(recordsAsInt []int) []string {
	length := len(recordsAsInt)
	recordsAsString := make([]string, length)
	for i := 0; i < length; i++ {
		valueAsString := strconv.Itoa(recordsAsInt[i])
		recordsAsString[i] = valueAsString
	}
	return recordsAsString
}
