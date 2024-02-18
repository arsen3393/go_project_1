package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"encoding/xml"
	"github.com/arthurkushman/go-hungarian"
)

func main() {
	data, err := ioutil.ReadFile("yakutsk.osm")
	if err != nil {
		log.Fatal(err)
	}

	// Распарсивание данных OSM
	var osm OSM
	if err := xml.Unmarshal(data, &osm); err != nil {
		log.Fatal(err)
	}

	orders, err := Parsing_orders("orders.csv")
	if err != nil{
		fmt.Errorf("Error parsing orders: %v", err)
		log.Fatal(err)
	}

	couriers, err := Parsing_couriers("couriers.csv")
	if err != nil{
		fmt.Errorf("Error parsing couriers: %v", err)
		log.Fatal(err)
	}

	costMatrix := make([][]float64, len(orders))
	for i := range orders {
		costMatrix[i] = make([]float64, len(couriers))
		for j := range couriers {
			pointA := osm.Nodes[orders[i].FROM_Node]
			pointB := osm.Nodes[couriers[j].Geo_Node]
			costMatrix[i][j] = Haversine(pointA.Lat, pointB.Lon, pointB.Lat, pointB.Lon) // Составляем матрицу стоимостей(в нашем случае это будут расстояния от курьера до заказа)
		}
	}

	result := hungarian.SolveMin(costMatrix)
	
	for orderID, courierMap := range result {
        for courierID, _ := range courierMap {
            fmt.Printf("Заказ №%d достался Курьеру №%d, стоимость заказа: %.3f рублей\n", orders[orderID].id, couriers[courierID].id, orders[orderID].Price)
        }
    }

}

