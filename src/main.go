package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"encoding/xml"
	"github.com/arthurkushman/go-hungarian"
)

func main() {
	data, err := ioutil.ReadFile("../data/yakutsk.osm")
	if err != nil {
		log.Fatal(err)
	}

	// Распарсивание данных OSM
	var osm OSM
	if err := xml.Unmarshal(data, &osm); err != nil {
		log.Fatal(err)
	}
	// Парсинг заказов
	orders, err := Parsing_orders("../data/orders.csv")
	if err != nil{
		fmt.Errorf("Error parsing orders: %v", err)
		log.Fatal(err)
	}
	// Парсинг курьеров
	couriers, err := Parsing_couriers("../data/couriers.csv")
	if err != nil{
		fmt.Errorf("Error parsing couriers: %v", err)
		log.Fatal(err)
	}
	dif := 0
	dif_couriers := 0
	// Заполнение структур для того чтобы матрица получилась квадратная
	if len(orders)<len(couriers){
		dif = len(couriers) - len(orders)
		for i:= 0; i < dif; i++{
			order := Order{id:404}
			orders = append(orders,order)
		}
	}else if len(orders)>len(couriers){
		dif_couriers = len(orders) - len(couriers)
		for i := 0; i < dif_couriers; i++{
			courier := Courier{id:404}
			couriers = append(couriers,courier)
		}
	}

	
	costMatrix := make([][]float64, len(orders)) 
	for i := range orders {
		costMatrix[i] = make([]float64, len(couriers))
		if (dif != 0) && (i >= len(orders)-1){
			for j := range couriers{
				costMatrix[i][j] = 999999
				
			} 
		} else {
			for j := range couriers {
				if (dif_couriers != 0) && (j >= len(couriers)-1){
					costMatrix[i][j] = 999999
				}else{
					pointA := osm.Nodes[orders[i].FROM_Node]
					pointB := osm.Nodes[couriers[j].Geo_Node]
					costMatrix[i][j] = Haversine(pointA.Lat, pointB.Lon, pointB.Lat, pointB.Lon)
				}
			} // Составляем матрицу стоимостей(в нашем случае это будут расстояния от курьера до заказа)		
		}
	}

	result := hungarian.SolveMin(costMatrix)

	for orderID, courierMap := range result { // вывод результата
        for courierID, _ := range courierMap {
			if (orders[orderID].id == 404){	
				fmt.Printf("Курьеру №%d пока ничего не досталось\n", couriers[courierID].id)
			}else if (couriers[courierID].id == 404){
				fmt.Printf("Заказ №%d пока без курьера\n", orders[orderID].id)
			}else{
            fmt.Printf("Заказ №%d достался Курьеру №%d, стоимость заказа: %.3f рублей\n", orders[orderID].id, couriers[courierID].id, orders[orderID].Price)
			}
        }
    }

}

