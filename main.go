package main

import (
	"embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type LocationId int

var data embed.FS

func main() {
	locationList1, locationList2 := loadLocationLists()

	listDifference := reconcileLists(locationList1, locationList2)
	fmt.Printf("Total distance of location lists is %d", listDifference)
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
