package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalCountInWeek := 0
	startDay := 7 * (week - 1)
	endDay := 7 * week

	for i := startDay; i < endDay; i++ {
		totalCountInWeek += birdsPerDay[i]
	}

	return totalCountInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
