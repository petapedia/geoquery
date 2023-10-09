package peda

import (
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func GetAllBangunanLineString(mongoconn *mongo.Database, collection string) []GeoJsonLineString {
	lokasi := atdb.GetAllDoc[[]GeoJsonLineString](mongoconn, collection)
	return lokasi
}
