package models

import (
	"backed/app/core/consts"
	"gorm.io/gorm"
)

// SendEmailCodeRequest 发送邮箱验证码请求参数
type SendEmailCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// RegisterRequest 用户注册请求参数
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Identity int64  `json:"identity"`
}

// LoginRequest 用户登录请求参数
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Identity int64  `json:"identity"`
}

// ChangeUserHeadRequest  获取用户请求头像
type ChangeUserHeadRequest struct {
	Email string `json:"email"`
	Head  string `json:"head"`
}

type UserInfo struct {
	gorm.Model
	Email    string `json:"email"`
	Head     string `json:"head"` // 头像
	Password string `json:"password"`
	Identity int64  `json:"identity"`
	Status   int64  `json:"status"`
}

// status 0 未审核 1 审核通过 2 审核未通过 3 账号停用

// NewUserInfo 创建一个新用户
func NewUserInfo(in RegisterRequest) *UserInfo {
	// 在这里做个判断 如果用户申请的是管理员账号 那么需要等待审核
	if in.Identity == 1 {
		return &UserInfo{
			Email:    in.Email,
			Password: in.Password,
			Identity: in.Identity,
			Head:     consts.DefaultHead, // 设置默认头像
			Status:   0,
		}
	}
	return &UserInfo{
		Email:    in.Email,
		Password: in.Password,
		Identity: in.Identity,
		Head:     consts.DefaultHead,
		Status:   1,
	}
}
