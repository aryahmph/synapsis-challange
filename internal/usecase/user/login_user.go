package user

import (
	"context"
	"errors"
)

var (
	ErrLoginUser_WrongCredential error = errors.New("LOGIN_USER.WRONG_CREDENTIAL")
)

func (u userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = u.passwordManager.CheckPasswordHash(password, user.PasswordHash)
	if err != nil {
		return "", ErrLoginUser_WrongCredential
	}

	return u.jwtManager.Generate(user.ID)
}
