package color

const (
	pre   = "\033["
	post  = ``
	reset = "\033[0m"
)

// Paint is a color to paint, either as a foreground or background paint
type Paint string

// Valid colors for ANSI terminals
const (
	BlackPaint      Paint = `0;30`
	DarkRedPaint    Paint = `0;31`
	DarkGreenPaint  Paint = `0;32`
	DarkYellowPaint Paint = `0;33`
	DarkBluePaint   Paint = `0;34`
	DarkPurplePaint Paint = `0;35`
	DarkCyanPaint   Paint = `0;36`
	LightGrayPaint  Paint = `0;37`

	DarkGrayPaint Paint = `1;30`
	RedPaint      Paint = `1;31`
	GreenPaint    Paint = `1;32`
	YellowPaint   Paint = `1;33`
	BluePaint     Paint = `1;34`
	PurplePaint   Paint = `1;35`
	CyanPaint     Paint = `1;36`
	WhitePaint    Paint = `1;37`

	nilPaint Paint = `nil`
)

// Brush is a function that let's you colorize strings directly.
type Brush func(string) string

// NewBrush gives you a brush that you can invoke directly to create colorized
// strings
func NewBrush(background, foreground Paint) Brush {
	return NewStyle(background, foreground).Brush()
}

// Style will give you colorized strings.  Styles are immutable.
type Style struct {
	bg   Paint
	fg   Paint
	code string
}

// NewStyle gives you a style ready to produce strings with the given
// background and foreground colors
func NewStyle(background, foreground Paint) Style {
	bg := background
	fg := foreground
	return Style{
		bg,
		fg,
		computeColorCode(bg, fg),
	}
}

// Brush is a function that can be used to color things directly, i.e:
//
//    red := NewStyle(BlackPaint, RedPaint).Brush()
//    fmt.Printf("This is %s\n", red("red"))
func (s Style) Brush() Brush {
	return func(text string) string {
		return s.code + text + reset
	}
}

// WithBackground copies the current style and return a new Style that
// has the desired background. The original Style is unchanged and you
// must capture the return value.
func (s Style) WithBackground(color Paint) Style {
	newS := s
	newS.bg = color
	newS.code = computeColorCode(newS.bg, newS.fg)
	return newS
}

// WithForeground copies the current style and return a new Style that
// has the desired foreground. The original Style is unchanged and you
// must capture the return value.
func (s Style) WithForeground(color Paint) Style {
	newS := s
	newS.fg = color
	newS.code = computeColorCode(newS.bg, newS.fg)
	return newS
}

func computeColorCode(bg, fg Paint) string {
	if bg == nilPaint {
		return pre + string(fg) + "m" + post
	}

	// The background code is the last color code prefixed by 4
	bgColor := bg[len(bg)-1]
	back := pre + "4" + string(bgColor) + "m" + post

	front := pre + string(fg) + "m" + post
	return back + front
}

// Monochrome

// Black gives black text on a white background
func Black() Brush { return NewBrush(WhitePaint, BlackPaint) }

// White gives white text on a dark gray background
func White() Brush { return NewBrush(DarkGrayPaint, WhitePaint) }

// LightGray gives light gray text on a black background
func LightGray() Brush { return NewBrush(nilPaint, LightGrayPaint) }

// Bright colors

// Blue gives blue text on a black background
func Blue() Brush { return NewBrush(nilPaint, BluePaint) }

// Cyan gives cyan text on a black background
func Cyan() Brush { return NewBrush(nilPaint, CyanPaint) }

// Green gives green text on a black background
func Green() Brush { return NewBrush(nilPaint, GreenPaint) }

// Purple gives purple text on a black background
func Purple() Brush { return NewBrush(nilPaint, PurplePaint) }

// Red gives red text on a black background
func Red() Brush { return NewBrush(nilPaint, RedPaint) }

// Yellow gives yellow text on a black background
func Yellow() Brush { return NewBrush(nilPaint, YellowPaint) }

// Dark colors

// DarkBlue gives dark blue text on a black background
func DarkBlue() Brush { return NewBrush(nilPaint, DarkBluePaint) }

// DarkCyan gives dark cyan text on a black background
func DarkCyan() Brush { return NewBrush(nilPaint, DarkCyanPaint) }

// DarkGray gives dark gray text on a black background
func DarkGray() Brush { return NewBrush(nilPaint, DarkGrayPaint) }

// DarkGreen gives dark green text on a black background
func DarkGreen() Brush { return NewBrush(nilPaint, DarkGreenPaint) }

// DarkPurple gives dark purple text on a black background
func DarkPurple() Brush { return NewBrush(nilPaint, DarkPurplePaint) }

// DarkRed gives dark red text on a black background
func DarkRed() Brush { return NewBrush(nilPaint, DarkRedPaint) }

// DarkYellow gives brown text on a black background
func DarkYellow() Brush { return NewBrush(nilPaint, DarkYellowPaint) }
