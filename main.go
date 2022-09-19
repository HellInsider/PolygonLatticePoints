package main

import (
	"fmt"
)

type Point struct {
	x, y, ind int
}

func main() {
	fmt.Println(LatticePoints(getPoints()))
}

func GCD(a, b int) int {
	for b != 0 {
		a %= b
		if a == 0 {
			return IntAbs(b)
		}
		b %= a
	}
	return IntAbs(a)
}

func IntAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func getPoints() (int, []Point) {
	var n, x, y int
	var points []Point
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		points = append(points, Point{x: x, y: y, ind: i + 1})
	}
	return n, points
}

func LatticePoints(n int, arr []Point) int {
	var points int

	margPoints := findPointsOnMargins(n, arr)
	gauss := findGauss(n, arr)

	points = gauss - margPoints + 2
	points /= 2

	return points
}

func findPointsOnMargins(n int, arr []Point) int {
	var sum int

	for i := 0; i < n-1; i++ {
		if arr[i].x-arr[i+1].x == 0 {
			sum += IntAbs(arr[i].y - arr[i+1].y)
		} else if arr[i].y-arr[i+1].y == 0 {
			sum += IntAbs(arr[i].x - arr[i+1].x)
		} else {
			sum += GCD(IntAbs(arr[i].x-arr[i+1].x), IntAbs(arr[i].y-arr[i+1].y))
		}
	}

	if arr[n-1].x-arr[0].x == 0 {
		sum += IntAbs(arr[n-1].y - arr[0].y)
	} else if arr[n-1].y-arr[0].y == 0 {
		sum += IntAbs(arr[n-1].x - arr[0].x)
	} else {
		sum += GCD(IntAbs(arr[n-1].x-arr[0].x), IntAbs(arr[n-1].y-arr[0].y))
	}

	//fmt.Println("Marg: ", sum)
	return sum
}

func findGauss(n int, arr []Point) int {
	var sum int
	for i := 1; i < n-1; i++ {
		sum += arr[i].x * (arr[i+1].y - arr[i-1].y)
		//fmt.Println(arr[i].x, "* (", arr[i+1].y, " - ", arr[i-1].y, ") = ", arr[i].x*(arr[i+1].y-arr[i-1].y))
	}

	sum += arr[n-1].x * (arr[0].y - arr[n-2].y)
	//fmt.Println(arr[n-1].x, "* (", arr[0].y, " - ", arr[n-2].y, ") = ", arr[n-1].x*(arr[0].y-arr[n-2].y))
	sum += arr[0].x * (arr[1].y - arr[n-1].y)
	//fmt.Println(arr[0].x, "* (", arr[1].y, " - ", arr[n-1].y, ") = ", arr[0].x*(arr[1].y-arr[n-1].y))

	//fmt.Println("Gauss: ", sum)
	return sum
}
