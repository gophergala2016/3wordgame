package validation

import (
	"regexp"
	"errors"
)

var msg_validator *regexp.Regexp = regexp.MustCompile(`^\S+\s+\S+\s+\S+\n$`)

var msg_formatter *regexp.Regexp = regexp.MustCompile(`\s{2,}`)

func ValidateMsg (msg string) (string, error) {
	if msg_validator.MatchString(msg) == true {
		return msg_formatter.ReplaceAllString(msg, " "), nil
	} else {
		return msg, errors.New("Invalid message")
	}
}
