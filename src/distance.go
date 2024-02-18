package main

import (
	"encoding/xml"
	"math"
)

type Node struct {
	Lat  float64 `xml:"lat,attr"`
	Lon  float64 `xml:"lon,attr"`
}

type OSM struct {
	XMLName xml.Name `xml:"osm"`
	Nodes   []Node   `xml:"node"`
}



// Формула гаверсинуса для вычисления расстояния
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371000 // Радиус Земли в метрах

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}

