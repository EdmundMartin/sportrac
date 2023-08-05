package gpx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadActivity(t *testing.T) {

	filename := "Strava.gpx"

	activity, err := LoadActivity(filename)
	assert.NoError(t, err)
	assert.NotNil(t, activity)
}

func TestLoadActivity_Metadata(t *testing.T) {
	activity := loadDefaultActivity(t)
	assert.Equal(t, "StravaGPX Android", activity.Creator)
	assert.Equal(t, "1.1", activity.Version)
}

func TestActivity_KilometerSplits(t *testing.T) {
	activity := loadDefaultActivity(t)
	splits := activity.KilometerSplits()
	assert.Equal(t, 7, len(splits))

	lastSplit := splits[len(splits)-1]
	assert.Equal(t, 0.713909812017795, lastSplit.Distance)
}

func TestActivity_CalculateAscentDescent(t *testing.T) {
	activity := loadDefaultActivity(t)

	elevation := activity.CalculateAscentDescent()
	assert.Equal(t, 18.000000000000007, elevation.Ascent)
	assert.Equal(t, 18.000000000000014, elevation.Descent)
}

func loadDefaultActivity(t *testing.T) *Activity {
	filename := "Strava.gpx"
	activity, err := LoadActivity(filename)
	require.NoError(t, err)
	require.NotNil(t, activity)
	return activity
}
