package main

import (
	"fmt"
	"lmrl/logic/bible"
	"log"
	"path/filepath"
)

const ProjectDir = "/Users/jimmy.jiang/workspace/bairex/lmrl"

func main() {
	// 加载文本数据
	bibleData, err := bible.LoadBibleData(filepath.Join(ProjectDir, "logic/bible/resources/bible.txt"))
	if err != nil {
		log.Fatalf("Failed to load bible text: %v", err)
	}

	// 保存为压缩的 protobuf
	if err := bible.SaveToCompressedProtobuf(bibleData,
		filepath.Join(ProjectDir, "logic/bible/resources/bible-data.pb.gz")); err != nil {
		log.Fatalf("Failed to save bible data: %v", err)
	}

	fmt.Println("Bible data processed successfully")
}
