// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user_access_log.sql

package model

import (
	"context"
	"database/sql"
	"time"
)

const listUserAccessLogByPagination = `-- name: ListUserAccessLogByPagination :many
select user_access_log.id, user_access_log.created_at, user_access_log.updated_at, user_access_log.deleted_at, user_access_log.user_id, user_access_log.ip, user_access_log.path, user_access_log.method, user_access_log.params, user_access_log.body, u.name as user_name
from user_access_log
         left join user u on u.id = user_access_log.user_id
limit ? offset ?
`

type ListUserAccessLogByPaginationParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListUserAccessLogByPaginationRow struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	UserID    uint64         `json:"user_id"`
	Ip        string         `json:"ip"`
	Path      string         `json:"path"`
	Method    string         `json:"method"`
	Params    string         `json:"params"`
	Body      string         `json:"body"`
	UserName  sql.NullString `json:"user_name"`
}

func (q *Queries) ListUserAccessLogByPagination(ctx context.Context, arg *ListUserAccessLogByPaginationParams) ([]*ListUserAccessLogByPaginationRow, error) {
	rows, err := q.query(ctx, q.listUserAccessLogByPaginationStmt, listUserAccessLogByPagination, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListUserAccessLogByPaginationRow
	for rows.Next() {
		var i ListUserAccessLogByPaginationRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.UserID,
			&i.Ip,
			&i.Path,
			&i.Method,
			&i.Params,
			&i.Body,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
