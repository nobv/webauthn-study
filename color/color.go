package color

import (
	"fmt"
)

type Color int

const (
	Brack Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
)

func Coloring(color Color, message string) string {

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, message)

}

func (color Color) String() string {
	switch color {
	case Brack:
		return "Brack"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Magenta:
		return "Magenta"
	case Cyan:
		return "Cyan"
	default:
		return ""
	}
}
