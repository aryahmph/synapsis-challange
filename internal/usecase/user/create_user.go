package user

import (
	"context"
	"errors"
	userModel "synapsis-challange/internal/model/user"
)

var (
	ErrCreateUser_UserExist error = errors.New("CREATE_USER.USER_EXIST")
)

func (u userUsecase) CreateUser(ctx context.Context, arg userModel.AddUser) (int32, error) {
	exist, err := u.userRepository.VerifyAvailableUser(ctx, arg.Email)
	if err != nil {
		return 0, err
	}

	if exist {
		return 0, ErrCreateUser_UserExist
	}

	password, err := u.passwordManager.HashPassword(arg.PasswordHash)
	if err != nil {
		return 0, err
	}
	arg.PasswordHash = password

	rid, err := u.userRepository.CreateUser(ctx, arg)
	if err != nil {
		return 0, err
	}
	return rid, nil
}
