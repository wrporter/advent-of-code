package color

// See https://misc.flogisoft.com/bash/tip_colors_and_formatting

var (
	Reset = "\033[0m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	BgBlack  = "\033[40m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgPurple = "\033[45m"
	BgCyan   = "\033[46m"
	BgGray   = "\033[47m"
	BgWhite  = "\033[107m"

	Bold       = "\033[1m"
	Dim        = "\033[2m"
	Underlined = "\033[4m"
	Blink      = "\033[5m"
	Inverted   = "\033[7m"
)
