package chinese

import (
	"embed"

	"github.com/golang/freetype/truetype"
	"github.com/golang/freetype"
)

//go:embed wqy-microhei.ttc
var chineseFontFS embed.FS

var chineseFont *truetype.Font

// GetFont returns the Chinese font (wqy-microhei).
// Font is loaded lazily on first call.
func GetFont() *truetype.Font {
	if chineseFont == nil {
		fontBytes, err := chineseFontFS.ReadFile("wqy-microhei.ttc")
		if err != nil {
			panic(err)
		}
		chineseFont, err = freetype.ParseFont(fontBytes)
		if err != nil {
			panic(err)
		}
	}
	return chineseFont
}
