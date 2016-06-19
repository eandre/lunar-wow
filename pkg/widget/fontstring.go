package widget

type JustifyH string
type JustifyV string

const (
	JustifyLeft   JustifyH = "LEFT"
	JustifyCenter JustifyH = "CENTER"
	JustifyRight  JustifyH = "RIGHT"

	JustifyTop    JustifyV = "TOP"
	JustifyMiddle JustifyV = "MIDDLE"
	JustifyBottom JustifyV = "BOTTOM"
)

type FontFlags string

const (
	FontMonochrome   FontFlags = "MONOCHROME"
	FontOutline      FontFlags = "OUTLINE"
	FontThickOutline FontFlags = "THICKOUTLINE"
)

type fontInstance interface {
	GetFont() (filename string, fontHeight float32, flags FontFlags)
	GetFontObject() Font
	GetJustifyH() JustifyH
	GetJustifyV() JustifyV
	GetShadowColor() (r, g, b, a float32)
	GetShadowOffset() (x, y float32)
	GetSpacing() float32
	GetTextColor() (r, g, b, a float32)

	SetFont(filename string, fontHeight float32, flags FontFlags)
	SetFontObject(obj interface{})
	SetJustifyH(justify JustifyH)
	SetJustifyV(justify JustifyV)
	SetShadowColor(r, g, b, a float32)
	SetShadowOffset(x, y float32)
	SetSpacing(spacing float32)
	SetTextColor(r, g, b, a float32)
}

type Font interface {
	fontInstance

	GetAlpha() float32
	SetAlpha(alpha float32)
}

type FontString interface {
	LayeredRegion
	fontInstance

	CanNonSpaceWrap() bool
	CanWordWrap() bool
	GetMaxLines() int
	GetNumLines() int
	GetStringHeight() float32
	GetStringWidth() float32
	GetText() string
	GetWrappedWidth() float32
	IsTruncated() bool
	SetFormattedText(fmt string, args ...interface{})
	SetMaxLines(num int)
	SetNonSpaceWrap(enable bool)
	SetText(text string)
	SetTextHeight(height float32)
	SetWordWrap(enable bool)
}
