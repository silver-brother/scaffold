package handler

type (
	UserListReq struct {
		Username string `form:"username" example:"张三"`                         // 用户名
		Nickname string `form:"nickname" example:"张三疯"`                        // 昵称
		Email    string `form:"email" example:"zhangsanfeng@163.com"`          // 邮箱
		Age      uint64 `form:"age" example:"101"`                             // 年龄
		Birthday string `form:"birthday" example:"2020-01-01"`                 // 生日
		Mobile   string `form:"mobile" binding:"mobile" example:"13888888888"` // 手机号
	}

	UserListItem struct {
		ID       uint64 `json:"id"`                          // 用户ID
		Username string `json:"username" example:"username"` // 用户名
		Nickname string `json:"nickname"`                    // 昵称
		Email    string `json:"email"`                       // 邮箱
		Password string `json:"password"`                    // 密码
	}
	UserListRes struct {
		Total int64           `json:"total"`
		List  []*UserListItem `json:"list"`
	}
)

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
