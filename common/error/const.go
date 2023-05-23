package error

import userUsecase "synapsis-challange/internal/usecase/user"

var (
	DomainErrorTranslatorDirectories = map[error]*ClientError{
		userUsecase.ErrCreateUser_UserExist:      NewInvariantError("CREATE_USER.USER_EXIST"),
		userUsecase.ErrLoginUser_WrongCredential: NewUnauthorizedError("LOGIN_USER.WRONG_CREDENTIAL"),
	}
)
