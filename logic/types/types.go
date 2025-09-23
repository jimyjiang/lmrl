package types

type FileName = string
type Sermon struct {
	Filename string // 文件名
	Title    string // 讲道主题
	Date     string // 日期，格式如"2025-09-10"
	Speaker  string // 讲员姓名
	Duration string // 时长，如"45分钟"
	FileSize string // 文件大小，如"12MB"
}

// var (
// 	MP3_DIR = "/root/web/灵命日粮"
// )
