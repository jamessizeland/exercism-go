package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	app := "default"
	for _, char := range log {
		if char == '❗' {
			app = "recommendation"
			break
		} else if char == '🔍' {
			app = "search"
			break
		} else if char == '☀' {
			app = "weather"
			break
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, rune := range runes {
		if rune == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// var newLog []rune
// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
