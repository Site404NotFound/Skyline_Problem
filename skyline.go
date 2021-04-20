/*
Time Complexity:
O(n log n)

How to compile:
go build skyline.go

How to run:
go build skyline.go
./skyline

Coding Problem:
We want to compute the skyline of a city given buildings positions and heights.

The input is a List of Buildings.
A Building (class provided) consist of the left side, the right side and the height.

The output is a List of SkylinePoints.
A SkylinePoint (class provided) consist of an x value and a height. The height corresponds
to the height of the skyline to the right of the x value until the next SkylinePoint.

Given a List of Buildings return a List of skyline points that represents the skyline
Modify the method computeSkyline() to contain your algorithm.

Example:
Input
Building(1, 3, 2) (Shown as A below)
Building(2, 3, 4) (Shown as B below)
Building(4, 6, 3) (Shown as C below)
Building(6, 7, 3) (Shown as D below)

Visual representation (Only for reference, no need to display)
     B
     B   C C D
   A A   C C D
   A A   C C D
 0 1 2 3 4 5 6 7

Output:
SkylinePoint(1, 2)
SkylinePoint(2, 4)
SkylinePoint(3, 0)
SkylinePoint(4, 3)
SkylinePoint(7, 0)
*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/guptarohit/asciigraph"
)

type Building struct {
	left   int
	right  int
	height int
}

type SkylinePoint struct {
	x      int
	height int
}

func main() {
	// List of different skylines to test.
	buildings := [][]Building{
		{
			{left: 1, right: 3, height: 2},
			{left: 2, right: 3, height: 4},
			{left: 4, right: 6, height: 3},
			{left: 6, right: 7, height: 3},
		},
		{
			{left: 3, right: 13, height: 14},
		},
		{
			{left: 4, right: 6, height: 3},
			{left: 1, right: 3, height: 2},
			{left: 6, right: 7, height: 3},
			{left: 2, right: 3, height: 4},
		},
		{
			{left: 0, right: 3, height: 10},
			{left: 0, right: 10, height: 5},
			{left: 8, right: 14, height: 6},
		},
	}
	// List of expected results for the skylines above.
	expectedSkylinePoints := [][]SkylinePoint{
		{
			{x: 1, height: 2},
			{x: 2, height: 4},
			{x: 3, height: 0},
			{x: 4, height: 3},
			{x: 7, height: 0},
		},
		{
			{x: 3, height: 14},
			{x: 13, height: 0},
		},
		{
			{x: 1, height: 2},
			{x: 2, height: 4},
			{x: 3, height: 0},
			{x: 4, height: 3},
			{x: 7, height: 0},
		},
		{
			{x: 0, height: 10},
			{x: 3, height: 5},
			{x: 8, height: 6},
			{x: 14, height: 0},
		},
	}
	// Loop through each skyline test and compute the critical points.
	for i, buildingsList := range buildings {
		skylinePointList := ComputeSkyline(buildingsList)
		DisplayResults(skylinePointList, expectedSkylinePoints[i])
	}
}

func ComputeSkyline(buildings []Building) []SkylinePoint {
	if len(buildings) == 0 {
		return []SkylinePoint{}
	}
	if len(buildings) == 1 {
		return []SkylinePoint{
			{x: buildings[0].left, height: buildings[0].height},
			{x: buildings[0].right, height: 0},
		}
	}
	midPoint := int(len(buildings) / 2)
	leftSkyline := ComputeSkyline(buildings[:midPoint])
	rightSkyline := ComputeSkyline(buildings[midPoint:])
	return Merge(leftSkyline, rightSkyline)
}

func Merge(left []SkylinePoint, right []SkylinePoint) []SkylinePoint {
	criticalPoints := []SkylinePoint{}
	leftHeight, rightHeight := 0, 0
	leftPointer, rightPointer := 0, 0
	cornerIndex := 0
	for leftPointer < len(left) && rightPointer < len(right) {
		if left[leftPointer].x < right[rightPointer].x {
			leftHeight = left[leftPointer].height
			cornerIndex = left[leftPointer].x
			leftPointer++
		} else if right[rightPointer].x < left[leftPointer].x {
			rightHeight = right[rightPointer].height
			cornerIndex = right[rightPointer].x
			rightPointer++
		} else {
			leftHeight = left[leftPointer].height
			rightHeight = right[rightPointer].height
			cornerIndex = right[rightPointer].x
			leftPointer++
			rightPointer++
		}
		maxHeight := Max(leftHeight, rightHeight)
		if len(criticalPoints) == 0 || criticalPoints[len(criticalPoints)-1].height != maxHeight {
			criticalPoints = append(criticalPoints, SkylinePoint{x: cornerIndex, height: maxHeight})
		}
	}
	criticalPoints = append(criticalPoints, right[rightPointer:]...)
	criticalPoints = append(criticalPoints, left[leftPointer:]...)
	return criticalPoints
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func DisplayResults(results []SkylinePoint, expected []SkylinePoint) {
	// Display the skyline results and confirm they match the expected outcome.
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Skyline Points Match: %v\n", strings.Title(
		strconv.FormatBool(reflect.DeepEqual(results, expected))))
	graphSkyline(results)
}

func graphSkyline(results []SkylinePoint) {
	// This part gets a little ugly.
	// Great graphing library but maybe not for this usecase.
	graphData := []float64{}
	if len(results) > 0 {
		currentIndex := 0
		currentPoint := results[currentIndex]
		currentHeight := 0
		for i := 0; i <= results[len(results)-1].x; i++ {
			if i == currentPoint.x {
				currentHeight = currentPoint.height
				if currentIndex < len(results)-1 {
					currentIndex++
					currentPoint = results[currentIndex]
				}
			}
			graphData = append(graphData, float64(currentHeight))
		}
	} else {
		for i := 0; i <= 15; i++ {
			graphData = append(graphData, 0)
		}
	}
	fmt.Printf("\nBuilding Skyline:\n%v\n", asciigraph.Plot(graphData))
}
