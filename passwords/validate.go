// Package passwords provides validations for passwords.
package passwords

import "errors"

var (
	errShort       = errors.New("password is too short (minimum is 8 characters)")
	errLong        = errors.New("password is too long (maximum is 128 characters)")
	errIllegalChar = errors.New("password can only contain ASCII printable characters except spaces")
)

// Validate validates password.
// If password is valid, this returns nil.
// In a default setting, minimum length is 8, maximum length is 128, and allowed code points are in the range [33, 126].
func Validate(password string) error {
	// TODO: Accept user's preference.
	switch l := len(password); {
	case l < 8:
		return errShort
	case l > 128:
		return errLong
	}
	for _, s := range password {
		if s < 33 || 126 < s {
			return errIllegalChar
		}
	}
	return nil
}
