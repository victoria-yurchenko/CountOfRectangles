package main

import (
	"fmt"
	"math/rand"
)

func sortByX(points [][]int, length int) [][]int {

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if points[0][i] > points[0][j] {
				temp := points[0][i]
				points[0][i] = points[0][j]
				points[0][j] = temp
			}
		}
	}

	return points
}

func countRectangles(points [][]int, length int) int {

	count := 0
	points = sortByX(points, length)

	for k := 0; k < length; k++ {

		// 0 - x, 1 - y coordinates
		currentX := points[0][k]
		currentY := points[1][k]

		for i := k + 1; i < length; i++ {
			for j := i + 1; j < length; j++ {

				if points[0][i] == currentX {
					count = getRectangleCount(points, i, length, currentY, count, 1, 0)
				}

				if points[1][i] == currentY {
					count = getRectangleCount(points, i, length, currentX, count, 0, 1)
				}
			}
		}
	}

	return count
}

func getRectangleCount(points [][]int, i int, length int, checkingCoordinate int, count int, index1 int, index2 int) int {

	currentCoordinate := points[index1][i]

	for j := i + 1; j < length; j++ {
		if points[index1][i] == checkingCoordinate {
			if contains(points, length, points[index2][i], currentCoordinate, index2, index1) {
				count++
			}
			break
		}
	}
	return count
}

func contains(points [][]int, length int, coordinate1 int, coordinate2 int, index1 int, index2 int) bool {

	for i := 0; i < length; i++ {
		if points[index1][i] == coordinate1 && points[index2][i] == coordinate2 {
			return true
		}
	}

	return false
}

func printPoints(points [][]int, length int) {

	fmt.Print("x: ")
	for j := 0; j < length; j++ {
		fmt.Print(points[0][j])
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Print("y: ")
	for j := 0; j < length; j++ {
		fmt.Print(points[1][j])
		fmt.Print(" ")
	}
	fmt.Println()
}

func main() {

	const countPoints = 2 // coordinates x, y
	length := rand.Intn(50)
	points := make([][]int, countPoints)

	for y := 0; y < countPoints; y++ {
		points[y] = make([]int, length)
	}

	for y := 0; y < countPoints; y++ {
		for x := 0; x < length; x++ {
			points[y][x] = rand.Intn(10)
		}
	}

	printPoints(points, length)

	fmt.Println(countRectangles(points, length))

}
