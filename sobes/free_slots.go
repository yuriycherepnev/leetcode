package main

import (
	"fmt"
	"sort"
)

func findFreeSlots(user1, user2 [][]float64, dayLimits []float64, duration float64) [][]float64 {
	allBusy := make([][]float64, 0, len(dayLimits))
	allBusy = append(allBusy, user1...)
	allBusy = append(allBusy, user2...)

	sort.Slice(allBusy, func(i, j int) bool {
		return allBusy[i][0] < allBusy[j][0]
	})

	// Объединяем пересекающиеся интервалы
	mergedBusy := [][]float64{allBusy[0]}

	for i := 1; i < len(allBusy); i++ {
		lastIdx := len(mergedBusy) - 1

		if allBusy[i][0] > mergedBusy[lastIdx][1] {
			mergedBusy = append(mergedBusy, allBusy[i])
		} else {
			if allBusy[i][1] > mergedBusy[lastIdx][1] {
				mergedBusy[lastIdx][1] = allBusy[i][1]
			}
		}
	}

	// Ищем свободные интервалы
	var freeSlots [][]float64
	currentTime := dayLimits[0]

	for _, busy := range mergedBusy {
		if busy[0]-currentTime >= duration {
			freeSlots = append(freeSlots, []float64{currentTime, busy[0]})
		}

		if busy[1] > currentTime {
			currentTime = busy[1]
		}
	}

	// Проверяем свободное время до конца рабочего дня
	if dayLimits[1]-currentTime >= duration {
		freeSlots = append(freeSlots, []float64{currentTime, dayLimits[1]})
	}

	return freeSlots
}

func main() {
	u1 := [][]float64{
		{9.0, 10.0},
		{12.0, 13.0},
		{16.0, 19.0},
	}

	u2 := [][]float64{
		{10.0, 11.5},
		{12.5, 14.5},
		{15.0, 16.5},
	}

	limits := []float64{9.0, 20.0}
	duration := 1.0

	result := findFreeSlots(u1, u2, limits, duration)

	fmt.Println("Свободные слоты:")
	for _, slot := range result {
		fmt.Printf("[%.1f, %.1f]\n", slot[0], slot[1])
	}
}
