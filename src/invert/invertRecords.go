package invert

import (
	"sync"
)

type pair struct {
	row, col int
}

func invert(pairs chan pair, records [][]int, inverted *[][]int, wg *sync.WaitGroup) {
	for pair := range pairs {
		(*inverted)[pair.col][pair.row] = records[pair.row][pair.col]
	}
	wg.Done()
}

func InvertRecords(records [][]int) [][]int {
	length := len(records)

	pairs := make(chan pair)
	var wg sync.WaitGroup
	wg.Add(length)

	inverted := make([][]int, length)
	for i := range inverted {
		inverted[i] = make([]int, length)
	}
	
	for i := 0; i < length; i++ {
		go invert(pairs, records, &inverted, &wg)
		
		for j := 0; j < length; j++ {
			pairs <- pair{row: i, col: j}
		}
	}

	close(pairs)
	wg.Wait()

	return inverted
}
