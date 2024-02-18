package main

import (
	"fmt"
	"github.com/arthurkushman/go-hungarian"
)

func main() {
	courierNames := []string{"Курьер 1", "Курьер 2", "Курьер 3", "Курьер 4", "Курьер 5", "Курьер 6", "Курьер 7", "Курьер 8"}
	// матрица стоимостей
	result := hungarian.SolveMin([][]float64{
		{6, 2, 3, 4, 5, 11, 3, 8},
		{3, 8, 2, 8, 1, 12, 5, 4},
		{7, 9, 5, 10, 2, 11, 6, 8},
		{6, 7, 3, 4, 3, 5, 5, 3},
		{1, 2, 6, 13, 9, 11, 3, 6},
		{6, 2, 3, 4, 5, 11, 3, 8},
		{4, 6, 8, 9, 7, 1, 5, 3},
		{9, 1, 2, 5, 2, 7, 3, 8},
	})

	// Вывод результата на экран
	for i, courierIndex := range result {
		
		for _, colIndex := range courierIndex {
			fmt.Println("Заказ", i+1, "назначен на", courierNames[int(colIndex)])
		}
	}
}