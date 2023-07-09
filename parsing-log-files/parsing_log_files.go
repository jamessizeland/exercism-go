package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]\s`)
	return re.MatchString(text)

}

func SplitLogLine(text string) []string {
	// split text with separator that is < > with any number of ~ or = in between
	re := regexp.MustCompile(`\<(~|=|\*|-)*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	// each line might contain one or more instances of the word password inside quotation marks. The quote might contain other words before or after the word password, and its letters might be in any case.
	re := regexp.MustCompile(`"(?i)[^"]*password[^"]*"`)
	count := 0
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		count += len(matches)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	// remove the string end-of-line followed by any number of digits
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// each line might contain a username in the format "User <username>"
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9_]+)`)

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			lines[i] = "[USR] " + matches[1] + " " + line
		}
	}
	return lines
}
