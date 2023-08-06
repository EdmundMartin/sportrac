package tcx

import "time"

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

// KilometerSplit represents a specific split for a particular kilometer within the recorded activity
type KilometerSplit struct {
	Time     time.Duration
	Start    time.Time
	End      time.Time
	Distance float64
}

func (tx *TcxFile) KilometerSplits() []KilometerSplit {
	var currentTotal float64
	var start Trackpoint
	var splits []KilometerSplit

	lap := tx.Activities.ActivityList[0].Lap.Track.Trackpoints
	for idx, trkPt := range lap {
		if idx == 0 {
			start = trkPt
			continue
		}
		newDistance := trkPt.DistanceMeters

		if currentTotal+newDistance >= 1_000 {
			splits = append(splits, KilometerSplit{
				Time:     trkPt.Time.Sub(start.Time),
				Start:    start.Time,
				End:      trkPt.Time,
				Distance: newDistance + currentTotal,
			})
			start = trkPt
			currentTotal = 0
		} else {
			currentTotal += newDistance
		}
	}
	if currentTotal > 0 {
		splits = append(splits, KilometerSplit{
			Time:     lap[len(lap)-1].Time.Sub(start.Time),
			Start:    start.Time,
			End:      lap[len(lap)-1].Time,
			Distance: currentTotal,
		})
	}

	return splits
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
