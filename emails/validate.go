// Package emails provides validations for emails.
package emails

import (
	"errors"
	"strings"
)

// Validate validates email.
// If email is valid, this returns nil.
func Validate(email string) error {
	// TODO: improve.
	l := len(email)

	if l == 0 {
		return errors.New("email cannot be blank")
	}
	if l > 254 {
		return errors.New("email is too long (maximum is 254 characters)")
	}

	i := strings.Index(email, "@")
	switch {
	case i < 0:
		return errors.New(`email should contain "@"`)
	case i == 0, i+1 == l, email[l-1] == '@':
		return errors.New(`email cannot start or end with "@"`)
	case i > 64:
		return errors.New(`there can be up to 64 characters long before "@" of email`)
	}

	i = strings.Index(email, ".")
	switch {
	case i < 0:
		return errors.New(`email must have some "."`)
	case i == 0, i+1 == l, email[l-1] == '.':
		return errors.New(`email cannot start or end with "."`)
	}

	switch {
	case strings.Index(email, " ") > 0:
		return errors.New(`email cannot contain " "`)
	case strings.Index(email, "..") > 0:
		return errors.New(`email cannot contain ".."`)
	case strings.Index(email, ".@") > 0:
		return errors.New(`email cannot have "." just before "@"`)
	}

	return nil
}
