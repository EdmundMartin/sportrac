package tcx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadFile(t *testing.T) {

	tcxF, err := LoadFile("Example.tcx")
	assert.NoError(t, err)
	assert.NotNil(t, tcxF)

	fmt.Println(tcxF)
}

func TestTcxFile_TotalDistance(t *testing.T) {

	txF := loadDefaultFile(t)

	result := txF.TotalDistance()
	assert.Equal(t, 9.7624433594, result)
}

func TestTcxFile_MaxHeartRate(t *testing.T) {
	txF := loadDefaultFile(t)
	result := txF.MaxHeartRate()
	assert.Equal(t, 194, result)
}

func loadDefaultFile(t *testing.T) *TcxFile {
	f, err := LoadFile("Example.tcx")
	require.NoError(t, err)
	require.NotNil(t, f)

	return f
}
