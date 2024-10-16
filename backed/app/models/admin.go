package models

// GetAllUserRequest 获取请求用户信息
type GetAllUserRequest struct {
	Page   int64 `json:"page"`
	Number int64 `json:"number"`
}

// UpdatePackageStatusRequest 更新包裹状态请求
type UpdatePackageStatusRequest struct {
	Email          string `json:"email"`
	PackageId      string `json:"package_id"`
	CurrentAddress string `json:"currentAddress"`
}
