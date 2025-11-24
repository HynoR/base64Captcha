package chinese

import (
	"image/color"
	"math/rand"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/mojocn/base64Captcha"
)

// Driver is a driver of unicode Chinese characters.
type Driver struct {
	//Height png height in pixel.
	Height int
	//Width Captcha png width in pixel.
	Width int

	//NoiseCount text noise count.
	NoiseCount int

	//ShowLineOptions := OptionShowHollowLine | OptionShowSlimeLine | OptionShowSineLine .
	ShowLineOptions int

	//Length random string length.
	Length int

	//Source is a unicode which is the rand string from.
	Source string

	//BgColor captcha image background color (optional)
	BgColor *color.RGBA

	//Fonts list of font names to use
	Fonts []string

	fontsArray []*truetype.Font
}

// NewDriver creates a driver of Chinese characters
func NewDriver(height int, width int, noiseCount int, showLineOptions int, length int, source string, bgColor *color.RGBA, fonts []string) *Driver {
	tfs := []*truetype.Font{}

	if len(fonts) > 0 {
		for _, fff := range fonts {
			tf := base64Captcha.DefaultEmbeddedFonts.LoadFontByName("fonts/" + fff)
			tfs = append(tfs, tf)
		}
	}

	// Always include Chinese font
	tfs = append(tfs, GetFont())

	return &Driver{
		Height:          height,
		Width:           width,
		NoiseCount:      noiseCount,
		ShowLineOptions: showLineOptions,
		Length:          length,
		Source:          source,
		BgColor:         bgColor,
		Fonts:           fonts,
		fontsArray:      tfs,
	}
}

// GenerateIdQuestionAnswer generates captcha content and its answer
func (d *Driver) GenerateIdQuestionAnswer() (id, content, answer string) {
	id = base64Captcha.RandomId()

	ss := strings.Split(d.Source, ",")
	length := len(ss)
	if length == 1 {
		c := base64Captcha.RandText(d.Length, ss[0])
		return id, c, c
	}
	if length <= d.Length {
		c := base64Captcha.RandText(d.Length, base64Captcha.TxtNumbers+base64Captcha.TxtAlphabet)
		return id, c, c
	}

	res := make([]string, d.Length)
	for k := range res {
		res[k] = ss[rand.Intn(length)]
	}

	content = strings.Join(res, "")
	return id, content, content
}

// DrawCaptcha generates captcha item(image)
func (d *Driver) DrawCaptcha(content string) (item base64Captcha.Item, err error) {
	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = base64Captcha.RandLightColor()
	}
	itemChar := base64Captcha.NewItemChar(d.Width, d.Height, bgc)

	//draw hollow line
	if d.ShowLineOptions&base64Captcha.OptionShowHollowLine == base64Captcha.OptionShowHollowLine {
		itemChar.DrawHollowLine()
	}

	//draw slime line
	if d.ShowLineOptions&base64Captcha.OptionShowSlimeLine == base64Captcha.OptionShowSlimeLine {
		itemChar.DrawSlimLine(3)
	}

	//draw sine line
	if d.ShowLineOptions&base64Captcha.OptionShowSineLine == base64Captcha.OptionShowSineLine {
		itemChar.DrawSineLine()
	}

	//draw noise
	if d.NoiseCount > 0 {
		source := base64Captcha.TxtNumbers + base64Captcha.TxtAlphabet + ",.[]<>"
		noise := base64Captcha.RandText(d.NoiseCount, strings.Repeat(source, d.NoiseCount))
		err = itemChar.DrawNoise(noise, d.fontsArray)
		if err != nil {
			return
		}
	}

	//draw content
	err = itemChar.DrawText(content, d.fontsArray)
	if err != nil {
		return
	}

	return itemChar, nil
}
