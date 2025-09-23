package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"lmrl/logic/bible"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

//go:embed SourceHanSansSC-VF.ttf
var fontData []byte

// 计算书籍完成度
func calculateCompletion(book *bible.Book) float64 {
	totalVerses := 0
	readVerses := 0

	// 跳过第0章（占位）
	for _, chapter := range book.Chapters[1:] {
		// 跳过第0节（占位）
		for _, verse := range chapter.Verses[0:] {
			totalVerses++
			// 检查是否已读（最后一个字符不是"*"）
			if len(verse) > 0 && verse[len(verse)-1] != '*' {
				readVerses++
			}
		}
	}

	if totalVerses == 0 {
		return 0.0
	}
	return float64(readVerses) / float64(totalVerses)
}

// 创建进度图
func createProgressChart(face font.Face, abbrs []string, completions []float64, outputPath string) error {
	const (
		rows          = 11
		cols          = 6
		cellSize      = 100
		padding       = 10
		headerHeight  = 30
		margin        = 20
		abbrFontSize  = 12
		percFontSize  = 10
		titleFontSize = 16
	)

	// 计算图像总尺寸
	width := cols*cellSize + (cols-1)*padding + 2*margin
	height := rows*cellSize + (rows-1)*padding + 2*margin + headerHeight

	// 创建图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充背景色
	bgColor := color.RGBA{240, 240, 240, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// 绘制标题
	title := "Bible Reading Progress"
	titleX := (width - len(title)*titleFontSize/2) / 2
	titleY := margin + titleFontSize
	drawString(face, img, title, titleX, titleY, titleFontSize, color.RGBA{0, 0, 0, 255})

	// 绘制格子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; {
			index := i*cols + j
			if index >= len(completions) {
				break
			}

			// 计算格子位置
			x := margin + j*(cellSize+padding)
			y := margin + headerHeight + i*(cellSize+padding)

			// 计算颜色（完成度越高颜色越深）
			completion := completions[index]
			green := uint8(50 + 200*completion)
			blue := uint8(200 - 150*completion)
			cellColor := color.RGBA{0, green, blue, 255}

			// 绘制格子背景
			drawRect(img, x, y, cellSize, cellSize, cellColor)

			// 绘制边框
			drawRectOutline(img, x, y, cellSize, cellSize, color.RGBA{200, 200, 200, 255})

			// 显示书籍缩写（假设abbreviations在另一个数组中）
			abbr := abbrs[index] // fmt.Sprintf("Book %d", index+1) // 替换为实际缩写
			abbrX := x + (cellSize-len(abbr)*abbrFontSize/2)/2
			abbrY := y + cellSize/2 - 10
			drawString(face, img, abbr, abbrX, abbrY, abbrFontSize, color.RGBA{0, 0, 0, 255})

			// 显示完成百分比
			percText := fmt.Sprintf("%.1f%%", completion*100)
			percX := x + (cellSize-len(percText)*percFontSize/2)/2
			percY := y + cellSize/2 + 10
			drawString(face, img, percText, percX, percY, percFontSize, color.RGBA{0, 0, 0, 255})

			j++
		}
	}

	// 保存图像
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// 辅助函数：绘制矩形
func drawRect(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	rect := image.Rect(x, y, x+width, y+height)
	draw.Draw(img, rect, &image.Uniform{c}, image.Point{}, draw.Src)
}

// 辅助函数：绘制矩形边框
func drawRectOutline(img *image.RGBA, x, y, width, height int, c color.RGBA) {
	for i := x; i < x+width; i++ {
		img.Set(i, y, c)
		img.Set(i, y+height-1, c)
	}
	for j := y; j < y+height; j++ {
		img.Set(x, j, c)
		img.Set(x+width-1, j, c)
	}
}

// 加载中文字体
func loadChineseFont() (font.Face, error) {
	f, err := opentype.Parse(fontData)
	if err != nil {
		return nil, fmt.Errorf("解析字体失败: %v", err)
	}

	// 创建字体Face，大小设为24
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("创建字体Face失败: %v", err)
	}

	return face, nil
}

// drawString 使用真实字体绘制文本
func drawString(face font.Face, img *image.RGBA, text string, x, y int, _ int, c color.Color) {
	// 创建字体绘制器
	drawer := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}

	// 绘制文本
	drawer.DrawString(text)
}

// 获取带日期的输出路径
func getOutputPath() (string, error) {
	// 创建.progress_charts目录（如果不存在）
	dir := "./progress_charts"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %v", err)
	}

	// 生成日期格式文件名
	now := time.Now()
	filename := fmt.Sprintf("%d%02d%02d.png", now.Year(), now.Month(), now.Day())

	return filepath.Join(dir, filename), nil
}

func main() {
	chineseFont, err := loadChineseFont()
	if err != nil {
		log.Fatalf("无法加载中文字体: %v", err)
	}
	// 获取输出路径
	outputPath, err := getOutputPath()
	if err != nil {
		log.Fatalf("获取输出路径失败: %v", err)
	}
	// 示例数据 - 实际应用中应该从BibleData结构获取
	bookmap := make(map[string]float64)
	completions := make([]float64, 66)
	abbrs := make([]string, 66)
	data, err := bible.LoadFromCompressedProtobuf()
	if err != nil {
		fmt.Print("Error loading bible data: " + err.Error())
	}
	for _, book := range data.Books {
		bookmap[book.Abbreviation] = calculateCompletion(book)
	}
	for i := 0; i < 66; i++ {
		abbr := bible.AbbrTable[i][1]
		abbrs[i] = abbr
		completions[i] = bookmap[abbr]
	}

	// 创建进度图
	if err := createProgressChart(chineseFont, abbrs, completions, outputPath); err != nil {
		fmt.Println("Error creating progress chart:", err)
		return
	}

	fmt.Println("Progress chart created successfully!")
}
