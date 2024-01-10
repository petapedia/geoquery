# GeoQuery

![image](https://github.com/petapedia/geoquery/assets/11188109/bfc085c6-e10f-438f-a528-c8ae7ee7da3b)

contoh query menggunakan geointersect :
```sh
db.petapedia.findOne({ geometry: { $geoIntersects: { $geometry: { type: "Point", coordinates: [ 107.57566599749208,-6.873806903286834 ] } } } })
db.petapedia.findOne({ geometry: { $geoIntersects: { $geometry: { type: "Point", coordinates: [ 107.57551038854763,-6.873444649526604 ] } } } })
db.petapedia.findOne({ geometry: { $geoIntersects: { $geometry: { type: "Point", coordinates: [ 107.57566599749208,-6.873806903286834 ] } } } })
```
contoh database nya :
```json
[{
  "_id": {
    "$oid": "6523daf0d4b7c6f836439258"
  },
  "type": "Feature",
  "properties": {
    "name": "Gedung Rektorat"
  },
  "geometry": {
    "coordinates": [
      [
        [
          107.57553607088482,
          -6.873686834063591
        ],
        [
          107.57550699801232,
          -6.873865265463579
        ],
        [
          107.57582151362982,
          -6.873930865225944
        ],
        [
          107.57585058650227,
          -6.87374980985922
        ],
        [
          107.57570522214144,
          -6.873720945953536
        ],
        [
          107.57553607088482,
          -6.873686834063591
        ]
      ]
    ],
    "type": "Polygon"
  }
},
{
  "_id": {
    "$oid": "6523dafed4b7c6f83643925a"
  },
  "type": "Feature",
  "properties": {
    "name": "Gedung Pendidikan"
  },
  "geometry": {
    "coordinates": [
      [
        [
          107.57525501601577,
          -6.87321358139036
        ],
        [
          107.5752283365386,
          -6.873583581896455
        ],
        [
          107.57578610435593,
          -6.873626624441528
        ],
        [
          107.57581611876782,
          -6.873261590420185
        ],
        [
          107.57525501601577,
          -6.87321358139036
        ]
      ]
    ],
    "type": "Polygon"
  }
},
{
  "_id": {
    "$oid": "6523db0ad4b7c6f83643925c"
  },
  "type": "Feature",
  "properties": {
    "name": "Auditorium"
  },
  "geometry": {
    "coordinates": [
      [
        107.57602973411036,
        -6.8732794028977935
      ],
      [
        107.57596954488866,
        -6.87365843069189
      ],
      [
        107.57622233962314,
        -6.8737028213150495
      ],
      [
        107.57628080915379,
        -6.873320378889417
      ],
      [
        107.57602973411036,
        -6.873281110230849
      ]
    ],
    "type": "LineString"
  }
}]
```

## ENV Setting
Peta Pedia Library Package

![image](https://github.com/petapedia/geoquery/assets/11188109/1009c159-7cb2-4d0f-bf78-5ad590a307ad)

## Release

```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.28                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/petapedia/geoquery@v0.0.28   #publish to pkg dev, replace ORG/URL with your repo URL
```
