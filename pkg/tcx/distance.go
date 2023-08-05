package tcx

func (tx *TcxFile) TotalDistance() float64 {
	var totalDistance float64
	for _, activity := range tx.Activities.ActivityList {
		// TODO - Can there be multiple laps per file
		totalDistance += activity.Lap.DistanceMeters / 1000
	}
	return totalDistance
}

func (tx *TcxFile) MaxHeartRate() int {
	var maxHeartRate int
	for _, activity := range tx.Activities.ActivityList {
		// TODO - Can there be multiple laps per file
		maxHeartRate = max(activity.Lap.MaximumHeartRateBpm.Value, maxHeartRate)
	}
	return maxHeartRate
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
