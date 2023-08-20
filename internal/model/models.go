// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UserSex string

const (
	UserSexF UserSex = "F"
	UserSexM UserSex = "M"
	UserSexO UserSex = "O"
)

func (e *UserSex) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserSex(s)
	case string:
		*e = UserSex(s)
	default:
		return fmt.Errorf("unsupported scan type for UserSex: %T", src)
	}
	return nil
}

type NullUserSex struct {
	UserSex UserSex `json:"user_sex"`
	Valid   bool    `json:"valid"` // Valid is true if UserSex is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserSex) Scan(value interface{}) error {
	if value == nil {
		ns.UserSex, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserSex.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserSex) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserSex), nil
}

// 用户信息表
type User struct {
	ID        uint64       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	// 姓名
	Name string `json:"name"`
	// 性别 1男 2女 3未知
	Sex UserSex `json:"sex"`
	// 生日
	BirthDate string `json:"birth_date"`
	// 身份证
	IDCard string `json:"id_card"`
	// 手机号
	Mobile string `json:"mobile"`
	Avatar string `json:"avatar"`
	// 简介
	Description string `json:"description"`
}

// 用户访问日志表
type UserAccessLog struct {
	ID uint64 `json:"id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 更新时间
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	// 用户id
	UserID uint64 `json:"user_id"`
	// ip地址
	Ip string `json:"ip"`
	// 访问路径
	Path string `json:"path"`
	// 请求方式
	Method string `json:"method"`
	// 请求参数
	Params string `json:"params"`
	// 请求体
	Body string `json:"body"`
}