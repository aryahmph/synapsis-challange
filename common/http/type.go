package http

type (
	Response struct {
		Data interface{} `json:"data"`
	}

	User struct {
		ID       int32  `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Address  string `json:"address"`
	}

	AddUser struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Address  string `json:"address" validate:"required"`
	}
	
	Login struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)
