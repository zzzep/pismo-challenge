package config

import (
	"regexp"
	"testing"
)

func TestGetDatabaseConnection(t *testing.T) {
	got := GetDatabaseConnection()
	pattern := `.*\:.*\@tcp\(.*:\d{1,5}\)\/.*\?.*`
	match, err := regexp.MatchString(pattern, got)
	if err != nil {
		t.Error("wrong pattern to connection: " + err.Error())
	}
	if !match {
		t.Error("wrong pattern to connection: " + got)
	}
}
