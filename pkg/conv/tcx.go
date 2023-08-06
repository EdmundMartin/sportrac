package conv

import (
	"encoding/xml"
	"github.com/EdmundMartin/sportrac/pkg/gpx"
	"github.com/EdmundMartin/sportrac/pkg/tcx"
	"os"
)

func TcxToGpx(tcxFile *tcx.TcxFile) *gpx.Activity {
	activity := &gpx.Activity{
		XMLName: xml.Name{
			Space: "http://www.topografix.com/GPX/1/1",
			Local: "gpx",
		},
		Metadata: gpx.ActivityMetaData{
			Time: tcxFile.Activities.ActivityList[0].Id,
		},
		Track: gpx.Track{
			Type:         tcxFile.Activities.ActivityList[0].Sport,
			TrackSegment: gpx.TrackSegment{TrackPoints: []gpx.TrackPoint{}},
		},
	}

	trackPoints := tcxFile.Activities.ActivityList[0].Lap.Track.Trackpoints
	for _, point := range trackPoints {
		activity.Track.TrackSegment.TrackPoints = append(activity.Track.TrackSegment.TrackPoints,
			gpx.TrackPoint{
				Longitude: point.Position.LongitudeDegrees,
				Latitude:  point.Position.LatitudeDegrees,
				Elevation: point.AltitudeMeters,
				Time:      point.Time,
			},
		)
	}
	return activity
}

func SaveGPX(filename string, activity *gpx.Activity) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	contents, err := xml.Marshal(activity)
	if err != nil {
		return err
	}
	if _, err := file.Write(contents); err != nil {
		return err
	}

	return nil
}
