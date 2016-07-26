package password

import "errors"

func Validate(password string) error {
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
