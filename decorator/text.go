package decorator

const (
	RED     = "red"
	GREEN   = "green"
	YELLOW  = "yellow"
	BLUE    = "blue"
	DEFAULT = "default"
)

func Colorize(text string, color string) string {
	switch color {
	case RED:
		return "\033[31m" + text + "\033[0m"
	case GREEN:
		return "\033[32m" + text + "\033[0m"
	case YELLOW:
		return "\033[33m" + text + "\033[0m"
	case BLUE:
		return "\033[34m" + text + "\033[0m"
	default:
		return text
	}
}
