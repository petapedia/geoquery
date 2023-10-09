package peda

import (
	"encoding/json"
)

func StringGetAllBangunan(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	datagedung := GetAllBangunanLineString(mconn, collectionname)
	jsondatagedung, _ := json.Marshal(datagedung)
	return string(jsondatagedung)
}
