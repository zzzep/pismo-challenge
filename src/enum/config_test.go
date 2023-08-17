package enum

import (
	"regexp"
	"testing"
)

func TestGetDatabaseConnection(t *testing.T) {
	got := GetDatabaseConnection()
	pattern := `.*\:.*\@tcp\(.*:\d{1,5}\)\/.*\?.*`
	match, _ := regexp.MatchString(pattern, got)
	if !match {
		t.Error("wrong pattern to connection: " + got)
	}
}
