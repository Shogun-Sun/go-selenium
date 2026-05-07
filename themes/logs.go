package themes

import (
	"fmt"
)

const (
	colorReset    = "\033[0m"
	colorRed      = "\033[31m"
	colorCritical = "\033[1;31m"
	colorGreen    = "\033[1;32m"
	colorBlue     = "\033[36m"
	colorGray     = "\033[90m"

	colorGrayBold = "\033[1;90m"
	colorAlert    = "\033[1;37;41m"
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

func LogTestCritical(title, format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("%s[CRITICAL] %s: %s%s\n", colorCritical, title, formattedText, colorReset)
}

func LogStep(format string, a ...any) {
	formattedText := fmt.Sprintf(format, a...)
	fmt.Printf("   %s │ %s%s\n", colorGrayBold, formattedText, colorReset)
}

func LogStartTest(name string) {
	LogTestInfo(fmt.Sprintf("%s Test", name), "Начат.")
}

func LogEndTest(name string) {
	LogTestPassed(fmt.Sprintf("%s Test", name), "Завершен.")
}

func RunTest(name string, testBody func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s[FATAL ERROR] %s: %v %s\n", colorAlert, name, r, colorReset)
		}
	}()

	LogStartTest(name)
	testBody()
	LogEndTest(name)
}
