/**
 * Места в кинотеатре расположены в один ряд. Только что пришедший зритель должен
 * чтобы сидеть максимально далеко от остальных зрителей в ряду.
 * То есть расстояние от того места, куда сядет зритель
 * до ближайшего к нему зрителя должно быть максимально.
 * Гарантируется, что в ряду всегда есть свободные места и уже сидит
 * хотя бы один зритель.
 * Напиши функцию, которая по заданному ряду мест (массиву из нулей и ед
 * вернёт расстояние от выбранного места до ближайшего зрителя.
 * Input: [1, 0, 0, 0, 1]
 * Output: 2
 * <p>
 * Input: [1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1]
 * Output: 2
 * <p>
 * Input: [1, 0, 1, 0]
 * Output: 1
 */

package main

import "fmt"

func solve(seats []int) int {
	left := -1
	maxDist := 0

	for right, v := range seats {
		if v != 0 {
			if left == -1 {
				maxDist = right
			} else {
				dist := (right - left) / 2
				if dist > maxDist {
					maxDist = dist
				}
			}
			left = right
		}
	}
	endDist := len(seats) - 1 - left
	if endDist > maxDist {
		maxDist = endDist
	}

	return maxDist
}

func main() {
	fmt.Println(solve([]int{0, 0, 0, 1, 0}))                   // 2
	fmt.Println(solve([]int{1, 0, 0, 0, 0, 1}))                // 2
	fmt.Println(solve([]int{1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1})) // 2
	fmt.Println(solve([]int{1, 0, 1, 0}))                      // 1
}

/*
без корнер кейсов

func solve(seats []int) int {
	left := -1
	maxDist := 0

	for right, v := range seats {
		if v != 0 {
			dist := (right - left) / 2
			if dist > maxDist {
				maxDist = dist
			}
			left = right
		}
	}

	return maxDist
}
*/
