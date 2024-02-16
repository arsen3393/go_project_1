package main

import(
	"fmt"
	//"github.com/RyanCarrier/dijkstra"
	//"log"
	"github.com/paulmach/osm"
	"github.com/fogleman/gg"
	"math"
)

const(
	Width = 1024
	Height = 1024
)

func drawMap(data *osm.OSM){
	dc := gg.NewContext(Width, Height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for _, n := range data.Nodes{
		x, y := latLonToXY(n.Lat, n.Lon)
		dc.DrawPoint(x, y, 1)
	}
	dc.SavePng("map.png")	
}

func latLonToXY(lat, lon float64) (float64, float64){
	rMajor := 6378137.0
	shift := math.Pi * rMajor
	x := lon * shift / 180
	y := math.Log(math.Tan((90+lat)*math.Pi/360)/(math.Pi/180))
	y = y * shift / 180
	return x, y
}

func main(){
	// graph:=dijkstra.NewGraph()
	// //Add the 3 verticies
	// graph.AddVertex(0)
	// graph.AddVertex(1)
	// graph.AddVertex(2)
	// //Add the arcs
	// graph.AddArc(0,1,1)
	// graph.AddArc(0,2,1)
	// graph.AddArc(1,0,1)
	// graph.AddArc(1,2,2)

	// best, err := graph.Shortest(0, 2)
	// 	if err!=nil{
	// 	log.Fatal(err)
	// 	}
	// fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
	//data, err := osm.DownloadMapData(61.9741, 129.5319, 129.5319, 129.9061)
	data, err := osm.DownloadMapData(61.9741, 129.5319, 62.0812, 129.9061)
	if err != nil{
		fmt.Println("Ошибка при загрузке данных карты", err)
		return
	}
	drawMap(data)
	
}
