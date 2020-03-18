package echo

import (
	"converter"
	"csv"
	"fmt"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csv.GetRecordsFromFileRequest(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		rowAsTring := converter.ConvertIntArrayToString(row)
		response = fmt.Sprintf("%s%s\n", response, strings.Join(rowAsTring, ","))
	}
	fmt.Fprint(w, response)
}
