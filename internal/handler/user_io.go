package handler

type (
	UserListReq struct {
		Pagination

		Username string `form:"username"` // 用户名
		Nickname string `form:"nickname"` // 昵称
		Email    string `form:"email"`    // 邮箱
		Age      uint64 `form:"age"`      // 年龄
		Birthday string `form:"birthday"` // 生日 2020-01-01
		Mobile   string `form:"mobile"`   // 手机号
	}

	UserListItem struct {
		Name        string `json:"name" binding:"required" example:"张三"`
		Sex         string `json:"sex" binding:"required,oneof=M F O" example="O"` // M: 男性 F: 女性 O: 其他
		BirthDate   string `json:"birth_date" binding:"required" example:"2020-01-01"`
		IDCard      string `json:"id_card" binding:"-" example:"123456789012345678"`
		Mobile      string `json:"mobile" binding:"-" example:"13888888888"`
		Avatar      string `json:"avatar" binding:"-" example:"https://www.baidu.com/img/bd_logo1.png"`
		Description string `json:"description" binding:"-" example:"这是一个用户"`
	}
	UserListRes struct {
		Total int64           `json:"total"`
		List  []*UserListItem `json:"list"`
	}
)

// UserHandle list method
type (
	UserCreateReq struct {
		Name        string `json:"name" binding:"required" example:"张三"`                                     // 姓名
		Sex         string `json:"sex" binding:"required,oneof=M F O" example:"O"`                             // 性别
		BirthDate   string `json:"birth_date" binding:"required" example:"2020-01-01"`                         // 生日
		CareType    uint64 `json:"care_type" binding:"required,oneof=1 2" example:"1"`                         // 1:身份证;2:护照
		IDCard      string `json:"id_card" binding:"required" example:"123456789012345678"`                    // 证件号码
		Mobile      string `json:"mobile" binding:"required" example:"1389999999"`                             // 手机号
		Avatar      string `json:"avatar" binding:"required" example:"https://www.baidu.com/img/bd_logo1.png"` // 头像地址
		Description string `json:"description" binding:"required" example:"这是一段人物的描述"`                // 描述
	}

	UserCreateRes struct {
	}
)
