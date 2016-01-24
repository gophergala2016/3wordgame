package validation

import (
	"regexp"
	"errors"
)

var msg_validator *regexp.Regexp = regexp.MustCompile(`^\S+\s+\S+\s+\S+\n$`)

var multiple_spaces *regexp.Regexp = regexp.MustCompile(`\s{2,}`)
var newline *regexp.Regexp = regexp.MustCompile(`\n$`)

func ValidateMsg (msg string) (string, error) {
	if msg_validator.MatchString(msg) == true {
		msg = multiple_spaces.ReplaceAllString(msg, " ")
		msg = newline.ReplaceAllString(msg, "")
		return msg, nil
	} else {
		return msg, errors.New("Invalid message")
	}
}

func StripNewLine(msg string) (string) {
	return newline.ReplaceAllString(msg, "")
}
