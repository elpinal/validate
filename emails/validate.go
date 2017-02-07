// Package emails provides validations for emails.
package emails

import (
	"errors"
	"strings"
)

var (
	errEmpty     = errors.New("email cannot be blank")
	errLong      = errors.New("email is too long (maximum is 254 characters)")
	errNoAt      = errors.New(`email should contain "@"`)
	errAt        = errors.New(`email cannot start or end with "@"`)
	errLongLocal = errors.New(`there can be up to 64 characters long before "@" of email`)
	errNoDot     = errors.New(`email must have some "."`)
	errDot       = errors.New(`email cannot start or end with "."`)
	errSpace     = errors.New(`email cannot contain " "`)
	errDotDot    = errors.New(`email cannot contain ".."`)
	errDotAt     = errors.New(`email cannot have "." just before "@"`)
)

// Validate validates email.
// If email is valid, this returns nil.
func Validate(email string) error {
	// TODO: improve.
	l := len(email)

	if l == 0 {
		return errEmpty
	}
	if l > 254 {
		return errLong
	}

	i := strings.Index(email, "@")
	switch {
	case i < 0:
		return errNoAt
	case i == 0, i+1 == l:
		return errAt
	case i > 64:
		return errLongLocal
	}

	i = strings.Index(email, ".")
	switch {
	case i < 0:
		return errNoDot
	case i == 0, i+1 == l:
		return errDot
	}

	switch {
	case strings.Index(email, " ") > 0:
		return errSpace
	case strings.Index(email, "..") > 0:
		return errDotDot
	case strings.Index(email, ".@") > 0:
		return errDotAt
	}

	return nil
}
