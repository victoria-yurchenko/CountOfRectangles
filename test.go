package main

import (
	"fmt"
	"math/rand"
)

func sortByX(points [][]int) [][]int {

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
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

	for l := 0; l < length; l++ {

		currentX := points[0][l]
		currentY := points[1][l]

		for i := l + 1; i < length; i++ {
			for j := i + 1; j < length; j++ {

				if points[0][i] == currentX {

					y2 := points[1][i]

					for k := i + 1; k < length; k++ {
						if points[1][i] == currentY {

							if contains(points, length, points[0][i], y2) {
								count++
							}
							break
						}
					}
				}

				if points[1][i] == currentY {

					x2 := points[0][i]

					for k := i + 1; k < length; k++ {
						if points[0][i] == currentX {

							if contains1(points, length, points[1][i], x2) {
								count++
							}
							break
						}
					}
				}
			}
		}
	}

	return count
}

func contains(points [][]int, length, x, y int) bool {

	for i := 0; i < length; i++ {
		if points[0][i] == x && points[1][i] == y {
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

func contains1(points [][]int, length, x, y int) bool {

	for i := 0; i < length; i++ {
		if points[1][i] == x && points[0][i] == y {
			return true
		}
	}

	return false
}

func main() {

	const countPoints = 2 // coordinates x, y
	//length := rand.Intn(20)
	length := 13
	points := make([][]int, countPoints)

	for y := 0; y < countPoints; y++ {
		points[y] = make([]int, length)
	}

	for y := 0; y < countPoints; y++ {
		for x := 0; x < length; x++ {
			points[y][x] = rand.Intn(10)
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if points[0][i] < points[0][j] {
				temp := points[0][i]
				points[0][i] = points[0][j]
				points[0][j] = temp
			}
		}
	}

	printPoints(points, length)

	fmt.Println(countRectangles(points, length))

}
