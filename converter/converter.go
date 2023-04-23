package converter

import "path/filepath"

type Converter struct {
	originalImg string
}

func NewConverter(imgPath string) *Converter {
	return &Converter{originalImg: imgPath}
}

func (c *Converter) Convert() {
	switch filepath.Ext(c.originalImg) {
	case ".jpg", ".JPG":
		c.JpgtoPng()
	case ".png", ".PNG":
		c.PngtoJpg()
	default:
		panic("error: invalid file extension")
	}
}
func (c *Converter) JpgtoPng() {

}

func (c *Converter) PngtoJpg() {

}
