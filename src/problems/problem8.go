package problems

import (
	"math"
	"strings"

	"github.com/MgShepherd/AdventOfCode2024/src/utils"
)

type Point struct {
	x int
	y int
}

func SolveProblem8() (int, error) {
	data, err := utils.ReadProblemFile(8)
	if err != nil {
		return 0, err
	}

	grid := utils.ConvertToGrid(data)
	locations := findAntennaLocations(grid)

	numUniqueLocations := getNumAntinodeLocations(grid, locations)
	return numUniqueLocations, nil
}

func findAntennaLocations(grid [][]string) map[string][]Point {
	locations := make(map[string][]Point)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != "." {
				newPoint := Point{x, y}
				if val, exists := locations[grid[y][x]]; exists {
					locations[grid[y][x]] = append(val, newPoint)
				} else {
					locations[grid[y][x]] = []Point{newPoint}
				}
			}
		}
	}

	return locations
}

func getNumAntinodeLocations(grid [][]string, locations map[string][]Point) int {
	numLocations := 0
	for _, v := range locations {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				points := getPointsFromAntennaLocations(v[i], v[j])
				for _, point := range points {
					addLocationIfUnique(&numLocations, grid, point)
				}
			}
		}
	}
	return numLocations
}

func getPointsFromAntennaLocations(location1, location2 Point) []Point {
	difference := Point{x: int(math.Abs(float64(location2.x - location1.x))),
		y: int(math.Abs(float64(location2.y - location1.y)))}

	points := []Point{}
	points = append(points,
		getPointFromDifference(difference, location1, location2),
		getPointFromDifference(difference, location2, location1))

	return points
}

func getPointFromDifference(difference, currentPoint, otherPoint Point) Point {
	newPoint := Point{x: -1, y: -1}

	if otherPoint.x < currentPoint.x {
		newPoint.x = currentPoint.x + difference.x
	} else {
		newPoint.x = currentPoint.x - difference.x
	}

	if otherPoint.y < currentPoint.y {
		newPoint.y = currentPoint.y + difference.y
	} else {
		newPoint.y = currentPoint.y - difference.y
	}

	return newPoint
}

func getDistanceFromPoint(current, starting Point) int {
	xDiff := int(math.Abs(float64(current.x - starting.x)))
	yDiff := int(math.Abs(float64(current.y - starting.y)))

	return xDiff + yDiff
}

func addLocationIfUnique(numLocations *int, grid [][]string, point Point) {
	if point.x >= 0 && point.x < len(grid[0]) && point.y >= 0 && point.y < len(grid) {
		if !strings.Contains(grid[point.y][point.x], "_") {
			*numLocations += 1
			grid[point.y][point.x] += "_"
		}
	}
}
