package validation

import (
	"testing"
)

func TestValidation(t *testing.T) {

	var msg string
	var err error

	_, err = ValidateMsg("one two\n")
	if err == nil {
		t.Errorf("one two \\n should've failed validation")
	}

	_, err = ValidateMsg("one two three four\n")
	if err == nil {
		t.Errorf("one two three four\\n should've failed validation")
	}

	msg, err = ValidateMsg("one two three\n")
	if err != nil {
		t.Errorf("one two three\\n should've passed validation")
	}

	if msg != "one two three" {
		t.Errorf("one two three\\n should've stayed that, but instead became %s", msg)
	}

	msg, err = ValidateMsg("one two  three\n")
	if err != nil {
		t.Errorf("one two  three\\n should've passed validation")
	}

	if msg != "one two three" {
		t.Errorf("one two  three\\n should've become one two three\\n, but instead became %s", msg)
	}

	msg, err = ValidateMsg("one. two  three\n")
	if err != nil {
		t.Errorf("one. two three\\n should've passed validation")
	}

	if msg != "one. two three" {
		t.Errorf("one. two three\\n should've stayed that, but instead became %s", msg)
	}
}
