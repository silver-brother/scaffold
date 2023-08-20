package handler

// UserHandle list method
type (
	UserCreateReq struct {
		Username string `json:"username" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Age      uint8  `json:"age" binding:"gte=1,lte=120"`
	}

	UserCreateRes struct {
	}
)
