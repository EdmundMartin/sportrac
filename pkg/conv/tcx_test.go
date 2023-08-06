package conv

import (
	"github.com/EdmundMartin/sportrac/pkg/tcx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestTcxToGpx(t *testing.T) {

	res, err := tcx.LoadFile("Example.tcx")
	require.NoError(t, err)
	require.NotNil(t, res)

	gpxResult := TcxToGpx(res)
	assert.Equal(t, 1213, len(gpxResult.Track.TrackSegment.TrackPoints))

	testFileName := "testFile"
	defer cleanFile(testFileName)

	err = SaveGPX(testFileName, gpxResult)
	require.NoError(t, err)
}

func cleanFile(target string) {
	os.Remove(target)
}