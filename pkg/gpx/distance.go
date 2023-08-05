package gpx

import (
	"fmt"
	"math"
	"time"
)

// KilometerSplit represents a specific split for a particular kilometer within the recorded activity
type KilometerSplit struct {
	Time     time.Duration
	Start    time.Time
	End      time.Time
	Distance float64
}

// PaceKm calculates the kilometer pace for a specific split
func (kms KilometerSplit) PaceKm() float64 {
	seconds := kms.Time.Seconds()
	minutesPerKm := 1.0 / ((kms.Distance / seconds) * 60)
	return 60 / minutesPerKm
}

// PaceMiles calculates the mile pace for a specific split
func (kms KilometerSplit) PaceMiles() float64 {
	return kms.PaceKm() * 0.621
}

func (kms KilometerSplit) String() string {
	return fmt.Sprintf("Time: %s Start:%s, End: %s", kms.Time.String(), kms.Start.String(), kms.End.String())
}

// TotalDistance returns the total distance covered during the tracked activity
func (a *Activity) TotalDistance() float64 {
	var lastPoint TrackPoint
	var totalKm float64
	for idx, point := range a.Track.TrackSegment.TrackPoints {
		if idx == 0 {
			lastPoint = point
			continue
		}
		totalKm += distanceBetweenPoints(lastPoint, point)
		lastPoint = point
	}
	return totalKm
}

func (a *Activity) TotalDistanceMiles() float64 {
	return a.TotalDistance() * 0.621
}

func (a *Activity) KilometerSplits() []KilometerSplit {
	var currentTotal float64
	var start TrackPoint
	var lastPoint TrackPoint
	var splits []KilometerSplit

	trackPoints := a.Track.TrackSegment.TrackPoints
	for idx, point := range trackPoints {
		if idx == 0 {
			lastPoint = point
			start = point
			continue
		}
		newDistance := distanceBetweenPoints(lastPoint, point)

		if newDistance+currentTotal >= 1 {
			splits = append(splits, KilometerSplit{
				Time:     point.Time.Sub(start.Time),
				Start:    start.Time,
				End:      point.Time,
				Distance: newDistance + currentTotal,
			})
			start = point
			currentTotal = 0
		} else {
			currentTotal += newDistance
		}
		lastPoint = point
	}
	if currentTotal > 0 {
		splits = append(splits, KilometerSplit{
			Time:     trackPoints[len(trackPoints)-1].Time.Sub(start.Time),
			Start:    start.Time,
			End:      trackPoints[len(trackPoints)-1].Time,
			Distance: currentTotal,
		})
	}

	return splits
}

func distanceBetweenPoints(first, second TrackPoint) float64 {
	longOne := degreesToRadian(first.Longitude)
	longTwo := degreesToRadian(second.Longitude)

	latOne := degreesToRadian(first.Latitude)
	latTwo := degreesToRadian(second.Latitude)

	distLon := longTwo - longOne
	distLat := latTwo - latOne

	a := math.Pow(math.Sin(distLat)/2, 2) + math.Cos(latOne)*math.Cos(latTwo)*math.Pow(math.Sin(distLon)/2, 2)

	c := 2 * math.Asin(math.Sqrt(a))

	var radius float64
	radius = 6371

	return c * radius
}

func degreesToRadian(value float64) float64 {
	return value * (math.Pi / 180)
}
