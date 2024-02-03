package user

// 创建用户请求结构体
type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// 创建用户响应
type CreateUserResponse struct {
	Username string `json:"username"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
}

// 修改密码请求
type ChangeRequest struct {
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// 修改响应
type ChangeResponse struct {
	NewPassword string `json:"new-password"`
}

// 修改用户请求
type ChangeNameRequest struct {
	Username string `json:"username"`
}

// 修改响应
type ChangeNameResponse struct {
	NewUserName string `json:"new-username"`
}
