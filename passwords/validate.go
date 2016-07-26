// Package passwords provides validations for passwords.
package passwords

import "errors"

// Validate validates password.
// If password is valid, this returns nil.
// In a default setting, minimum length is 8, maximum length is 128, and allowed code points are in the range [33, 126].
func Validate(password string) error {
	// TODO: Support specifying settings.
	if l := len(password); l < 8 {
		return errors.New("password is too short (minimum is 8 characters)")
	} else if l > 128 {
		return errors.New("password is too long (maximum is 128 characters)")
	}

	for _, s := range password {
		if s < 33 || 126 < s {
			return errors.New("password can only contain ASCII printable characters except spaces")
		}
	}
	return nil
}
