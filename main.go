package geoquery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/petapedia/geoquery/config"
	"github.com/petapedia/geoquery/helper"
	"github.com/petapedia/geoquery/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == config.EndpointSecret {
		resp.Response = GetLokasi(config.Mongocon, msg.Longitude, msg.Latitude)
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, helper.Json(resp))
}

func GetLokasi(mongoconn *mongo.Database, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("lokasi")
	filter := bson.M{
		"batas": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
			},
		},
	}
	var lokasi models.Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		fmt.Printf("GetLokasi: %v\n", err)
	}
	return lokasi.Nama

}
