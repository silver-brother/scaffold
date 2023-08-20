package handler

import (
	"scaffold/common/errorx"
	"scaffold/common/httpx"
	"scaffold/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	svcCtx *service.ServiceContext
}

func NewUserHandle(svcCtx *service.ServiceContext) *UserHandle {
	return &UserHandle{
		svcCtx: svcCtx,
	}
}

// List user list
// @Summary user list
// @Description get user list
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} UserCreateRes "返回用户信息"
// @Router /user [get]
func (l *UserHandle) List(ctx *gin.Context) {
	user, err := l.svcCtx.DBClient.GetUserById(ctx, 1)
	exists, err := errorx.CheckDBErr(err)
	if err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error500(ctx)
		return
	}
	if !exists {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msgf("user not found, id: %d", 1)
		httpx.Error404(ctx)
		return
	}
	httpx.Success(ctx, user)
}

// Create 新增用户
// @Summary 新增用户
// @Description 新增用户
// @Produce json
// @Param body body UserCreateReq true "body参数"
// @Success 200 {object} UserCreateRes "返回用户信息"
// @Router /user [post]
func (l *UserHandle) Create(ctx *gin.Context) {
	var req UserCreateReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error400(ctx, err.Error())
		return
	}

	httpx.Success(ctx, nil)
}
