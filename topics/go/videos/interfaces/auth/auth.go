package auth

type AuthError struct {
	User   string
	Reason string
}

func (e *AuthError) Error() string {
	return e.Reason
}

type User struct{}

func Login(user, passwd string) (*User, error) {
	var err error

	// TODO

	return nil, err
}
