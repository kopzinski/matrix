package csv

import (
	"converter"
	"encoding/csv"
	"net/http"
)

func GetRecordsFromFileRequest(r *http.Request) ([][]int, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	recordsAsInt, err := converter.ConvertStringMatrixToInt(records)
	if err != nil {
		return nil, err
	}
	return recordsAsInt, nil
}
