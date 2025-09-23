package jobs

import (
	"fmt"
	"lmrl/logic/cache"
	"lmrl/logic/types"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func SetupTest() {
	types.MP3_DIR = "testdata/mp3files"
}
func TestRunDownloadMp3Job(t *testing.T) {
	SetupTest()
	dt, err := time.Parse(time.DateOnly, "2026-01-05")
	require.NoError(t, err)
	filename := generateFileName(dt)
	require.EqualValues(t, "mw260105.mp3", filename)
	err = RunDownloadMp3Job(dt)
	require.NoError(t, err)
	maxFiles = 22
	dt, err = time.Parse(time.DateOnly, "2026-01-01")
	require.NoError(t, err)
	err = RunDownloadMp3Job(dt)
	require.NoError(t, err)

	// dt, err = time.Parse(time.DateOnly, "2025-09-25")
	// require.NoError(t, err)
	// err = RunDownloadMp3Job(dt)
	// require.NoError(t, err)
}

func TestRebuildMp3Cache(t *testing.T) {
	SetupTest()
	err := rebuildMp3Cache()
	require.NoError(t, err)
	fmt.Printf("mp3cache: %s\n", cache.GetMp3Cache())
}
