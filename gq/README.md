# Geospatioal Query Operators

<https://www.mongodb.com/docs/manual/reference/operator/query-geospatial/>

## GeoIntersects

![image](https://github.com/petapedia/geoquery/assets/11188109/55d0346d-f731-4116-a594-8d7bbfa553c3)

Untuk mengetahui apakah sebuah titik lokasi beririsan dengan sebuah poligon yang mana dari database.
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

## GeoWithin

Pada dokumen mongo :

```json
{
   <location field>: {
      $geoWithin: {
         $geometry: {
            type: <"Polygon" or "MultiPolygon"> ,
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
       $geoWithin: {
          $geometry: {
             type: "Polygon" ,
             coordinates: [
          [
            [
              107.57563124089006,
              -6.87337660729348
            ],
            [
              107.57563124089006,
              -6.8735235494584686
            ],
            [
              107.5757273342715,
              -6.8735235494584686
            ],
            [
              107.5757273342715,
              -6.87337660729348
            ],
            [
              107.57563124089006,
              -6.87337660729348
            ]
          ]
        ]
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
   "$geoWithin": bson.M{
    "$geometry": bson.M{
     "type":        "Polygon",
     "coordinates": [][]float64{long, lat},
    },
   },
  },
 }
```

## Near

Mencari lokasi terdekat dari titik kordinat.
Pada dokumen mongo :

```json
{
   <location field>: {
     $near: {
       $geometry: {
          type: "Point" ,
          coordinates: [ <longitude> , <latitude> ]
       },
       $maxDistance: <distance in meters>,
       $minDistance: <distance in meters>
     }
   }
}
```

Query mongosh :

```sh
db.lokasi.find(
   {
     batas: {
       $near: {
          $geometry: {
             type: "Point" ,
             coordinates: [ 107.57569081895566, -6.8735086203166285 ]
          },
          $maxDistance: 25,
          $minDistance: 0
       }
     }
   }
)
```

Implementasi Golang :

```go
filter := bson.M{
  "batas": bson.M{
   "$near": bson.M{
    "$geometry": bson.M{
     "type":        "Point",
     "coordinates": []float64{long, lat},
    },
    "$maxDistance": 25,
    "$minDistance": 0,
   },
  },
 }
```

## NearSphere

Mencari lokasi terdekat dari titik kordinat.
Pada dokumen mongo :

```json
{
  $nearSphere: {
     $geometry: {
        type : "Point",
        coordinates : [ <longitude>, <latitude> ]
     },
     $minDistance: <distance in meters>,
     $maxDistance: <distance in meters>
  }
}
```

Query mongosh :

```sh
db.lokasi.find(
   {
     batas: {
       $nearSphere: {
          $geometry: {
             type: "Point" ,
             coordinates: [ 107.57569081895566, -6.8735086203166285 ]
          },
          $maxDistance: 25,
          $minDistance: 0
       }
     }
   }
)
```

Implementasi Golang :

```go
filter := bson.M{
  "batas": bson.M{
   "$near": bson.M{
    "$geometry": bson.M{
     "type":        "Point",
     "coordinates": []float64{long, lat},
    },
    "$maxDistance": 25,
    "$minDistance": 0,
   },
  },
 }
```

## Box

Mencari lokasi di dalam kotak.
Pada dokumen mongo :

```json
{
  <location field>: {
     $geoWithin: {
        $box: [
          [ <bottom left coordinates> ],
          [ <upper right coordinates> ]
        ]
     }
  }
}
```

Query mongosh :

```sh
db.lokasi.find(
   {
     batas: {
       $geoWithin: {
          $box: [
                [ 107.56777632469857, -6.876878154969793 ],
                [ 107.58746411897147, -6.866868253263269 ]
             ]
       }
     }
   }
)
```
