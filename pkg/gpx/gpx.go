package gpx

import (
	"encoding/xml"
	"os"
	"time"
)

type Activity struct {
	XMLName  xml.Name         `xml:"gpx"`
	Creator  string           `xml:"creator,attr"`
	Version  string           `xml:"version,attr"`
	Metadata ActivityMetaData `xml:"metadata"`
	Track    Track            `xml:"trk"`
}

type ActivityMetaData struct {
	Time time.Time `xml:"time"`
}

type Track struct {
	Name         string       `xml:"name"`
	Type         string       `xml:"type"`
	TrackSegment TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	TrackPoints []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Longitude float64   `xml:"lon,attr"`
	Latitude  float64   `xml:"lat,attr"`
	Elevation float64   `xml:"ele"`
	Time      time.Time `xml:"time"`
}

func LoadActivity(filename string) (*Activity, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var activity Activity
	err = xml.Unmarshal(contents, &activity)
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func abs(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}
