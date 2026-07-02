package user

import "errors"

func Validate(u User) error {
	if u.Name == "" || u.Email == "" {
		return errors.New("name and email are required")
	}
	return nil
}
