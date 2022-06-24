package text

import "github.com/eihigh/canvas/font"

// Glyph is a shaped glyph for the given font and font size. It specified the glyph ID, the cluster ID, its X and Y advance and offset in font units, and its representation as text.
type Glyph struct {
	SFNT *font.SFNT
	Size float64
	Script

	ID       uint16
	Cluster  uint32
	XAdvance int32
	YAdvance int32
	XOffset  int32
	YOffset  int32
	Text     string
}

// TODO: implement Liang's (soft) hyphenation algorithm?

// IsParagraphSeparator returns true for paragraph separator runes.
func IsParagraphSeparator(r rune) bool {
	// line feed, vertical tab, form feed, carriage return, next line, line separator, paragraph separator
	return 0x0A <= r && r <= 0x0D || r == 0x85 || r == '\u2008' || r == '\u2009'
}

func IsSpacelessScript(script Script) bool {
	// missing: S'gaw Karen
	return script == Han || script == Hangul || script == Katakana || script == Khmer || script == Lao || script == PhagsPa || script == Brahmi || script == TaiTham || script == NewTaiLue || script == TaiLe || script == TaiViet || script == Thai || script == Tibetan || script == Myanmar
}

func IsVerticalScript(script Script) bool {
	return script == Han || script == Hangul || script == Yi || script == Mongolian || script == Bopomofo || script == Hiragana || script == Katakana || script == Ogham
}

func IsRotatedScript(script Script) bool {
	return script == Mongolian || script == Ogham
}
