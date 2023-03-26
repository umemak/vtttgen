package vtttgen

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
)

type Info struct {
	idx   int
	from  string
	to    string
	fName string
	x     int
	y     int
	w     int
	h     int
}

func Run(ctx *cli.Context) error {
	files, err := getTargetFiles(ctx.String("target"), "jpg")
	if err != nil {
		return fmt.Errorf("getTargetFiles: %w", err)
	}
	ih, iw, err := getImageSize(files[0])
	if err != nil {
		return fmt.Errorf("getImageSize: %w", err)
	}
	infos := build(ih, iw, ctx.Int("rows"), ctx.Int("columns"), files)
	print(infos)
	return nil
}

func getTargetFiles(target, ext string) ([]string, error) {
	pattern := filepath.Join(target, "*."+ext)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("filepath.Glob: %w", err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("%s not found", pattern)
	}
	return files, nil
}

func getImageSize(fn string) (int, int, error) {
	file, err := os.Open(fn)
	if err != nil {
		return 0, 0, fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, fmt.Errorf("image.Decode: %w", err)
	}
	return img.Bounds().Dy(), img.Bounds().Dx(), nil
}

func timeString(sec int) string {
	return time.Date(0, 0, 0, 0, 0, sec, 0, time.Local).Format("15:04:05.000")
}

func build(ih, iw, rows, cols int, files []string) []Info {
	h := ih / rows
	w := iw / cols
	i := 0
	infos := []Info{}
	for _, fn := range files {
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				idx := (rows * cols * i) + row*cols + col
				infos = append(infos, Info{
					idx:   idx + 1,
					from:  timeString(idx),
					to:    timeString(idx + 1),
					fName: filepath.Base(fn),
					x:     col * w,
					y:     row * h,
					w:     w,
					h:     h,
				})
			}
		}
		i++
	}
	return infos
}

func print(infos []Info) {
	fmt.Println("WEBVTT")
	for _, v := range infos {
		v.print()
	}
}

func (i Info) print() {
	fmt.Printf("\n%d\n", i.idx)
	fmt.Printf("%s --> %s\n", i.from, i.to)
	fmt.Printf("%s#xywh=%d,%d,%d,%d\n", i.fName, i.x, i.y, i.w, i.h)
}
