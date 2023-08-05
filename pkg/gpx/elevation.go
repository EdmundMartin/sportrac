package gpx

import "math"

// ElevationData holds total ascent and descent during a specific activity
type ElevationData struct {
	Ascent  float64
	Descent float64
}

// MaxElevation records the max elevation reached at any time during an activity
func (a *Activity) MaxElevation() float64 {
	max := float64(math.MinInt32)
	for _, point := range a.Track.TrackSegment.TrackPoints {
		if point.Elevation > max {
			max = point.Elevation
		}
	}
	return max
}

// MinElevation records the min elevation reached at any time during an activity
func (a *Activity) MinElevation() float64 {
	min := math.MaxFloat64
	for _, point := range a.Track.TrackSegment.TrackPoints {
		if point.Elevation < min {
			min = point.Elevation
		}
	}
	return min
}

// CalculateAscentDescent calculates ascent/descent totals
func (a *Activity) CalculateAscentDescent() *ElevationData {
	ev := &ElevationData{}

	var prev float64
	for idx, point := range a.Track.TrackSegment.TrackPoints {
		if idx == 0 {
			prev = point.Elevation
			continue
		}
		diff := prev - point.Elevation
		if diff < 0 {
			ev.Descent += abs(diff)
		} else {
			ev.Ascent += abs(diff)
		}
		prev = point.Elevation
	}
	return ev
}
