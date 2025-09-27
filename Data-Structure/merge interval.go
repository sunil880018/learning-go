package main

import (
	"fmt"
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Sort intervals based on the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		lastMerged := merged[len(merged)-1]

		if currentInterval[0] <= lastMerged[1] {
			// Merge intervals by updating the end time of the last merged interval
			lastMerged[1] = int(math.Max(float64(lastMerged[1]), float64(currentInterval[1])))
		} else {
			// If they don't overlap, add the current interval to the result
			merged = append(merged, currentInterval)
		}
	}

	return merged
}

func main() {
	// Example usage:
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergedIntervals := merge(intervals)
	fmt.Println("Merged Intervals:", mergedIntervals)
}
