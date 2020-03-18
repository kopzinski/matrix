package multiply

import (
	"csv"
	"fmt"
	"net/http"
	"strconv"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {

	records, err := csv.GetRecordsFromFileRequest(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string

	result := MultiplyValues(records)
	response = fmt.Sprintf("%s%s\n", response, strconv.Itoa(result))

	fmt.Fprint(w, response)
}
