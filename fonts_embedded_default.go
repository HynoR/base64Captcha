package base64Captcha

import "embed"

// defaultEmbeddedFontsFS Built-in font storage (non-Chinese fonts only).
// For Chinese font support, import the chinese subpackage.
//
//go:embed fonts/*.ttf
var defaultEmbeddedFontsFS embed.FS

var DefaultEmbeddedFonts = NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)
