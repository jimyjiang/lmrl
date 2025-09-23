package mp3file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParseMp3File(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取主目录:", err)
		return
	}
	s, err := ParseMP3File(filepath.Join(home, "doc/基督/灵命日粮/202601/mw260105.mp3"))
	if err != nil {
		t.Errorf("ParseMP3File() error = %v", err)
	}
	t.Logf("ParseMP3File() = %v", s)
}
