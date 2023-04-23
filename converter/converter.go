package converter

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	originalImgs []string
}

func NewConverter(imgPaths []string) *Converter {
	return &Converter{originalImgs: imgPaths}
}

func (c *Converter) Convert() {
	for _, img := range c.originalImgs {
		switch filepath.Ext(img) {
		case ".jpg", ".JPG":
			c.JpgtoPng(img)
		case ".png", ".PNG":
			c.PngtoJpg(img)
		default:
			panic("error: invalid file extension")
		}
	}
}

func (c *Converter) JpgtoPng(orig string) {
	// 元画像を開く
	f, err := os.Open(orig)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 元画像をデコードする
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	// 出力ファイルを作成する
	outDir := filepath.Dir(orig)
	outPath := filepath.Join(outDir, filepath.Base(orig[:len(orig)-len(filepath.Ext(orig))]+".png"))
	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// PNG形式でエンコードして出力する
	if err := png.Encode(out, img); err != nil {
		panic(err)
	}
}

func (c *Converter) PngtoJpg(orig string) {
	// 元画像を開く
	f, err := os.Open(orig)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 元画像をデコードする
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	// 出力ファイルを作成する
	outDir := filepath.Dir(orig)
	outPath := filepath.Join(outDir, filepath.Base(orig[:len(orig)-len(filepath.Ext(orig))]+".png"))
	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// JPEG形式でエンコードして出力する
	if err := jpeg.Encode(out, img, nil); err != nil {
		panic(err)
	}
}
