package main

import (
	"context"
	"fmt"
	"lmrl/logic/mp3file"
	"lmrl/logic/types"
	"log"
	"os"
	"os/signal"
	"path"
	"sort"
	"strings"
)

const (
	MP3_DIR = "/Users/jimmy.jiang/doc/基督/灵命日粮/卡1"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer func() {
		stop()
	}()
	mp3file.StartWorker(ctx)
	mapSermons, err := mp3file.GetSermonsFromDir(MP3_DIR, false)
	if err != nil {
		log.Fatalf("get sermons from dir error: %v", err)
	}
	sermons := make([]*types.Sermon, 0, len(mapSermons))
	for _, sermon := range mapSermons {
		sermons = append(sermons, sermon)
	}
	// 按日期倒序排列
	sort.Slice(sermons, func(i, j int) bool {
		return sermons[i].Date < sermons[j].Date
	})
	// 生成字符串slice， index+1 文件名 标题名
	var content []string
	for index, sermon := range sermons {
		content = append(content, fmt.Sprintf("%d. %s", index+1, sermon.Title))
	}
	// 写文件
	err = os.WriteFile(path.Join(MP3_DIR, "sermons.txt"),
		[]byte(fmt.Sprintf("%s\n", strings.Join(content, "\n"))),
		0644)
	if err != nil {
		log.Fatalf("write sermons to file error: %v", err)
	}
}
