// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package model

import (
	"context"
)

type Querier interface {
	GetUserById(ctx context.Context, id uint64) (*User, error)
	GetUserByIdCard(ctx context.Context, idCard string) (*User, error)
	InsertUser(ctx context.Context, arg *InsertUserParams) error
	ListUserAccessLogByPagination(ctx context.Context, arg *ListUserAccessLogByPaginationParams) ([]*ListUserAccessLogByPaginationRow, error)
	ListUserByIds(ctx context.Context, ids []uint64) ([]*User, error)
}

var _ Querier = (*Queries)(nil)
