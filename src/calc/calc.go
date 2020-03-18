package calc

import (
	"sync"
)

type pair struct {
	row, col int
}

type Operation func(pairs chan pair, records [][]int, result *int, wg *sync.WaitGroup)

func Multiply(pairs chan pair, records [][]int, result *int, wg *sync.WaitGroup) {
	if *result == 0 {
		*result = 1
	}
	for pair := range pairs {
		*result = *result * records[pair.row][pair.col]
	}
	wg.Done()
}

func Sum(pairs chan pair, records [][]int, result *int, wg *sync.WaitGroup) {
	for pair := range pairs {
		*result = *result + records[pair.row][pair.col]
	}
	wg.Done()
}

func BulkCalculate(records [][]int, fn Operation) int {
	result := 0
	length := len(records)

	pairs := make(chan pair)
	var wg sync.WaitGroup
	wg.Add(length)

	for i := 0; i < length; i++ {
		go fn(pairs, records, &result, &wg)
		for j := 0; j < length; j++ {
			pairs <- pair{row: i, col: j}
		}
	}
	close(pairs)
	wg.Wait()

	return result

}
