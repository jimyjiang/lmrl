package types

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	MP3_DIR = ""
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取主目录:", err)
		return
	}
	MP3_DIR = filepath.Join(home, "doc/基督/灵命日粮/202510")
}
