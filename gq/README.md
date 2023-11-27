# Geospatioal Query Operators

## GeoIntersects

Pada dokumen mongo :

```json
{
  <location field>: {
     $geoIntersects: {
        $geometry: {
           type: "<GeoJSON object type>" ,
           coordinates: [ <coordinates> ]
        }
     }
  }
}
```

Query mongosh :

```sh
db.lokasi.find(
   {
     batas: {
       $geoIntersects: {
          $geometry: {
             type: "Point" ,
             coordinates: [ 107.57569081895566, -6.8735086203166285 ]
          }
       }
     }
   }
)
```

Implementasi Golang :

```go
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
```
