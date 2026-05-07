package themes

import (
	"fmt"
)

const (
	colorReset    = "\033[0m"
	colorRed      = "\033[31m"
	colorGreen    = "\033[1;32m"
	colorBlue     = "\033[36m"
	colorGray     = "\033[90m"
	colorGrayBold = "\033[1;90m"
)

func LogTestPassed(title, format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("%s[PASS] %s: %s%s\n", colorGreen, title, formattedText, colorReset)
}

func LogTestInfo(title, format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("%s[INFO] %s: %s%s\n", colorBlue, title, formattedText, colorReset)
}

func LogTestError(title, format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("%s[ERROR] %s: %s%s\n", colorRed, title, formattedText, colorReset)
}

func LogStep(format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("   %s │ %s%s\n", colorGrayBold, formattedText, colorReset)
}
