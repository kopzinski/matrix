package flatten

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
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csv.GetRecordsFromFileRequest(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string

	flatten := FlattenRecords(records)

	flattenAsTring := converter.ConvertIntArrayToString(flatten)
	response = fmt.Sprintf("%s%s\n", response, strings.Join(flattenAsTring, ","))
	fmt.Fprint(w, response)
}
