package main

import (
	"fmt"
	"lmrl/logic/bible"
	"log"
	"os"
	"path/filepath"
)

const ProjectDir = "workspace/bairex/lmrl"

func main() {
	// 加载文本数据
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取主目录:", err)
		return
	}

	bibleData, err := bible.LoadBibleData(filepath.Join(home, ProjectDir, "logic/bible/resources/bible.txt"))
	if err != nil {
		log.Fatalf("Failed to load bible text: %v", err)
	}

	// 保存为压缩的 protobuf
	if err := bible.SaveToCompressedProtobuf(bibleData,
		filepath.Join(home, ProjectDir, "logic/bible/resources/bible-data.pb.gz")); err != nil {
		log.Fatalf("Failed to save bible data: %v", err)
	}

	fmt.Println("Bible data processed successfully")
}
