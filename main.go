package main

import (
	"embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type LocationId = uint

var data embed.FS

func main() {
	locationList1, locationList2 := loadLocationLists()

	listDifference := reconcileLists(locationList1, locationList2)
	fmt.Printf("Total distance of location lists is %d\n", listDifference)
	similarityScore := calculateSimilarityScore(locationList1, locationList2)
	fmt.Printf("Similarity score is %d\n", similarityScore)

	reports := loadReports()
	safeReports := filterSafeReports(reports)
	fmt.Printf("Total safe reports %d\n", safeReports)
}

func loadLocationLists() ([]LocationId, []LocationId) {
	lines, rErr := readLines("./data/location_input.txt")
	if rErr != nil {
		panic(rErr)
	}
	list1 := make([]LocationId, len(lines))
	list2 := make([]LocationId, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		v1, _ := strconv.Atoi(parts[0])
		v2, _ := strconv.Atoi(parts[1])
		list1[i] = LocationId(v1)
		list2[i] = LocationId(v2)
	}
	return list1, list2
}

func reconcileLists(list1 []LocationId, list2 []LocationId) uint {
	slices.Sort(list1)
	slices.Sort(list2)
	var totalDiffs uint = 0
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			totalDiffs = totalDiffs + uint(list1[i]-list2[i])
		} else {
			totalDiffs = totalDiffs + uint(list2[i]-list1[i])
		}
	}
	return totalDiffs
}

func calculateSimilarityScore(list1 []LocationId, list2 []LocationId) uint {
	freqMap := make(map[LocationId]uint)
	for _, value := range list2 {
		if count, ok := freqMap[value]; ok {
			count++
			freqMap[value] = count
		} else {
			freqMap[value] = 1
		}
	}
	simScore := uint(0)
	for _, value := range list1 {
		if count, ok := freqMap[value]; ok {
			simScore += value * count
		}
	}
	return simScore
}

func loadReports() []string {
	reports, rErr := readLines("./data/report_input.txt")
	if rErr != nil {
		panic(rErr)
	}
	return reports
}

type LevelDirection = int

var (
	LevelDirectionUnset      LevelDirection = 2
	LevelDirectionIncreasing                = 1
	LevelDirectionDecreasing                = 0
)

func filterSafeReports(reports []string) int {
	safeCount := 0
	for _, report := range reports {
		direction := LevelDirectionUnset
		safe := true
		levels := strings.Split(report, " ")
		for i := 0; i < len(levels)-1; i++ {
			curr, _ := strconv.Atoi(levels[i])
			next, _ := strconv.Atoi(levels[i+1])
			diff := curr - next
			absDiff := math.Abs(float64(diff))
			if absDiff < 1 || absDiff > 3 {
				safe = false
				break
			}
			var diffDirection LevelDirection
			if diff > 0 {
				diffDirection = LevelDirectionIncreasing
			} else {
				diffDirection = LevelDirectionDecreasing
			}
			if direction != LevelDirectionUnset {
				if diffDirection != direction {
					safe = false
					break
				}
			} else {
				direction = diffDirection
			}
		}
		if safe {
			safeCount++
		}
	}
	return safeCount
}
