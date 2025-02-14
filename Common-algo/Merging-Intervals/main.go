package main

import (
	"fmt"
	"sort"
)

// Interval struct
type Interval struct {
	Start int
	End   int
}

// MergeIntervals merges overlapping intervals
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]

		if curr.Start <= last.End { // Overlapping
			if curr.End > last.End {
				last.End = curr.End
			}
		} else { // No overlap, add new interval
			merged = append(merged, curr)
		}
	}

	return merged
}

func main() {
	intervals := []Interval{
		{1, 3}, {3, 6}, {5, 7}, {7, 10},
	}

	//	intervals := []Interval{
	//	{1, 3}, {2, 6}, {8, 10}, {15, 18},
	//}
	fmt.Println(intervals)
	merged := MergeIntervals(intervals)

	for _, interval := range merged {
		fmt.Printf("[%d, %d] ", interval.Start, interval.End)
	}
	fmt.Println("")
}
