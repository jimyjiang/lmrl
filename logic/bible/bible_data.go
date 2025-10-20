package bible

import (
	"bytes"
	"compress/gzip"
	"embed"
	"google.golang.org/protobuf/proto"
	"io"
	"io/fs"
)

//go:embed all:resources/*
var embeddedResources embed.FS

// LoadFromCompressedProtobuf 从压缩的 protobuf 文件加载圣经数据
// 直接返回 protobuf 生成的 BibleData 对象
func LoadFromCompressedProtobuf() (*BibleData, error) {
	// 读取压缩文件
	compressed, err := fs.ReadFile(embeddedResources, "resources/bible-data.pb.gz")
	if err != nil {
		return nil, err
	}

	// 解压数据
	gzReader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	buffer, err := io.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}

	// 解码 protobuf 消息
	var bibleData BibleData
	if err := proto.Unmarshal(buffer, &bibleData); err != nil {
		return nil, err
	}

	return &bibleData, nil
}
