package util

import (
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// LoadTTF loads a font given a path & size.
func LoadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeError := file.Close(); err != nil {
			panic(closeError)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	loadedFont, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(loadedFont, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
