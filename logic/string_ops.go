package logic

import (
	"bufio"
	"os"
	"strings"
)

func NewlineSeparate(lines string) string{
	return strings.NewReplacer(
			"\n", " ",
		).Replace(lines)
}

func PeriodNewLine(lines string) string{
	lines = strings.NewReplacer(
		". ", ".",
		).Replace(lines)

	return strings.NewReplacer(".", ".\n").Replace(lines)
}

func StrStdin() string{
	endCondition := true
	counter := 0

	var stringInput string

	for endCondition {
		var stringLine string
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		stringLine = scanner.Text()

		stringInput += stringLine + "\n"

		if stringLine == "" {
			counter += 1
		} else {
			counter = 0
		}

		if counter == 3{
			endCondition = false
		}
	}
	return stringInput
}
