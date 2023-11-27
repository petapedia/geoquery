package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, DataStuct any) {
	jsondata, _ := json.Marshal(DataStuct)
	fmt.Fprintf(w, string(jsondata))
}
