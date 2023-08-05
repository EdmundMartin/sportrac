package tcx

import (
	"encoding/xml"
	"os"
	"time"
)

type TcxFile struct {
	XMLName    xml.Name   `xml:"TrainingCenterDatabase"`
	Activities Activities `xml:"Activities"`
}

type Activities struct {
	ActivityList []Activity `xml:"Activity"`
}

type Activity struct {
	Id    time.Time `xml:"Id"`
	Sport string    `xml:"Sport,attr"`
	Lap   Lap       `xml:"Lap"`
}

type Lap struct {
	StartTime           time.Time      `xml:"StartTime,attr"`
	TotalTimeSeconds    float64        `xml:"TotalTimeSeconds"`
	DistanceMeters      float64        `xml:"DistanceMeters"`
	MaximumSpeed        float64        `xml:"MaximumSpeed"`
	Calories            int            `xml:"Calories"`
	AverageHeartRateBpm HeartRateValue `xml:"AverageHeartRateBpm"`
	MaximumHeartRateBpm HeartRateValue `xml:"MaximumHeartRateBpm"`
	Intensity           string         `xml:"Intensity"`
	Cadence             string         `xml:"Cadence"`
	TriggerMethod       string         `xml:"TriggerMethod"`
	Track               Track          `xml:"Track"`
}

type Track struct {
	Trackpoints []Trackpoint `xml:"Trackpoint"`
}

type Trackpoint struct {
	Time           time.Time      `xml:"Time"`
	HeartRateBpm   HeartRateValue `xml:"HeartRateBpm"`
	AltitudeMeters float64        `xml:"AltitudeMeters"`
	DistanceMeters float64        `xml:"DistanceMeters"`
	SensorState    string         `xml:"SensorState"`
	Position       Position       `xml:"Position"`
}

type Position struct {
	LatitudeDegrees  float64 `xml:"LatitudeDegrees"`
	LongitudeDegrees float64 `xml:"LongitudeDegrees"`
}

type HeartRateValue struct {
	Value int `xml:"Value"`
}

func LoadFile(filename string) (*TcxFile, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var file TcxFile

	err = xml.Unmarshal(contents, &file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
