package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"fmt"
)
   
type Order struct{
	FROM_Node int // точка откуда значение от 0 до 39637
	TO_Node int // точка куда значение от 0 до 39637
	Price float64 // цена заказа
	id int // ид заказа

}

type Courier struct{
	Geo_Node int // геолокация курьера значение от 0 до 39637
	id int // ид курьера
}

func Parsing_orders(filename string) ([]Order, error) { //функция парсинга заказов и csv
    ordersFile, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("Error opening orders file: %v", err)
    }
    defer ordersFile.Close()

    orderData := csv.NewReader(ordersFile)
    orderRecords, err := orderData.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("Error reading orders file: %v", err)
    }

    var orders []Order
    for _, record := range orderRecords {
        fromNode, err := strconv.Atoi(record[0])
        if err != nil {
            return nil, fmt.Errorf("Error converting order's FROM_Node: %v", err)
        }

        toNode, err := strconv.Atoi(record[1])
        if err != nil {
            return nil, fmt.Errorf("Error converting order's TO_Node: %v", err)
        }

        price, err := strconv.ParseFloat(record[2], 64)
        if err != nil {
            return nil, fmt.Errorf("Error converting order's Price: %v", err)
        }

        id, err := strconv.Atoi(record[3])
        if err != nil {
            return nil, fmt.Errorf("Error converting order's id: %v", err)
        }

        order := Order{FROM_Node: fromNode, TO_Node: toNode, Price: price, id: id}
        orders = append(orders, order)
    }

    return orders, nil
}

func Parsing_couriers(filename string) ([]Courier, error){ //функция парсинга заказов и csv
	couriersFile, err := os.Open(filename)
	if err != nil{
		return nil, fmt.Errorf("Error opening couriersfile: %v", err)
	}
	courierData := csv.NewReader(couriersFile)
	courierRecords, err := courierData.ReadAll()
	if err != nil{
		return nil, fmt.Errorf("Error reading couriersfile: %v", err)
	}
	var couriers []Courier
	for _, record := range courierRecords{
		geo_Node, err := strconv.Atoi(record[0])
		if err != nil{
			return nil, fmt.Errorf("Error converting courier's id")
		}
		id, err := strconv.Atoi(record[1])
		if err != nil{
			return nil, fmt.Errorf("Error converting courier's id")
		}
		courier := Courier{Geo_Node: geo_Node, id:id}
		couriers = append(couriers, courier)
	}

	return couriers, nil

}