package handler

import (
	"net/http"
	"scaffold/common/httpx"
	"scaffold/internal/model"
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

// List 用户列表
// @Summary 用户列表
// @Description 用户列表
// @Tags user
// @Accept json
// @Produce json
// @Param queryList query UserListReq false "查询参数"
// @Success 200 {object} httpx.Resp{data=UserListRes} "响应结果"
// @Router /v1/user [get]
func (l *UserHandle) List(ctx *gin.Context) {
	// 参数验证
	var req UserListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error200(ctx, http.StatusBadRequest, err.Error())
		return
	}

	limit, offset := httpx.Pagination(req.Page, req.PageSize)

	// 数据查询
	user, err := l.svcCtx.DBClient.ListUserByPagination(ctx, &model.ListUserByPaginationParams{Limit: limit, Offset: offset})
	if err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error500(ctx)
		return
	}

	list := make([]*UserListItem, 0)
	for _, item := range user {
		list = append(list, &UserListItem{
			Name:        item.Name,
			Sex:         string(item.Sex),
			BirthDate:   item.BirthDate,
			IDCard:      item.IDCard,
			Mobile:      item.Mobile,
			Avatar:      item.Avatar,
			Description: item.Description,
		})
	}
	httpx.Success(ctx, UserListRes{
		Total: 0,
		List:  list,
	})
}

// Create 新增用户
// @Summary 新增用户
// @Description 新增用户
// @Tags user
// @Accept json
// @Produce json
// @Param payload body UserCreateReq true "body参数"
// @Success 200 {object} httpx.Resp "响应结果"
// @Router /v1/user [post]
func (l *UserHandle) Create(ctx *gin.Context) {
	// 参数验证
	var req UserCreateReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error400(ctx, err.Error())
		return
	}

	err = l.svcCtx.DBClient.InsertUser(ctx, &model.InsertUserParams{
		Name:        req.Name,
		Sex:         model.UserSex(req.Sex),
		BirthDate:   req.BirthDate,
		IDCard:      req.IDCard,
		Mobile:      req.Mobile,
		Avatar:      req.Avatar,
		Description: req.Description,
	})
	if err != nil {
		l.svcCtx.Log.Error().Ctx(ctx.Request.Context()).Msg(err.Error())
		httpx.Error500(ctx)
		return
	}

	httpx.Success(ctx, nil)
}
