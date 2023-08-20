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

// @Summary 用户列表
// @Description 用户列表
// @Tags user
// @Accept json
// @Produce json
// @Param id query int false "用户ID"
// @Success 200 {object} httpx.Resp{} "响应结果"
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

// @Summary 新增用户
// @Description 新增用户
// @Tags 分组名称
// @Accept json
// @Produce json
// @Param body body UserCreateReq true "body参数"
// @Success 200 {object} httpx.Resp "响应结果"
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
