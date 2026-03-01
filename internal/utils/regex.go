package utils

import "regexp"

var logStartPattern = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3}) - (INFO|WARN|ERROR|DEBUG|TRACE)\s+\[`)

func IsNewEntry(line string) bool {
	return logStartPattern.MatchString(line)
}
