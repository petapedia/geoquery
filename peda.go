package peda

import (
	"encoding/json"
	"net/http"
)

func GCFPeda(w http.ResponseWriter, r *http.Request) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	datagedung := GetAllBangunanLineString(mconn, "petapedia")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datagedung)

	//fmt.Fprintf(w, datagedung)
}
