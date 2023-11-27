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
