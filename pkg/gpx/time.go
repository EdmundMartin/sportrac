package gpx

import "time"

func (a *Activity) TotalTime() time.Duration {
	size := len(a.Track.TrackSegment.TrackPoints)
	return a.Track.TrackSegment.TrackPoints[size-1].Time.Sub(a.Track.TrackSegment.TrackPoints[0].Time)
}
