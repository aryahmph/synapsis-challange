package user

type (
	User struct {
		ID           int32
		Email        string
		PasswordHash string
		Name         string
		Address      string
	}

	AddUser struct {
		Email        string
		PasswordHash string
		Name         string
		Address      string
	}
)
