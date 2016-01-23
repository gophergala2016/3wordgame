package validation

import (
	"regexp"
	"errors"
)

var msg_validator *regexp.Regexp = regexp.MustCompile(`^\w+\W+\w+\W+\w+\n$`)

var msg_formatter *regexp.Regexp = regexp.MustCompile(`\W{2,}`)

func ValidateMsg (msg string) (string, error) {
	if msg_validator.MatchString(msg) == true {
		return msg_formatter.ReplaceAllString(msg, " "), nil
	} else {
		return msg, errors.New("Invalid message")
	}
}
