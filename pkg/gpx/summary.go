package gpx

import "time"

type Summary struct {
	TotalDistance float64
	TotalTime     time.Duration
	Splits        []KilometerSplit
	Elevation     ElevationData
}

func (a *Activity) ProduceSummary() Summary {
	summary := Summary{}
	summary.TotalDistance = a.TotalDistance()
	summary.TotalTime = a.TotalTime()
	summary.Splits = a.KilometerSplits()
	summary.Elevation = *a.CalculateAscentDescent()
	return summary
}
