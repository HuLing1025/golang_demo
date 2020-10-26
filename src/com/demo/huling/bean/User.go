package bean

// 用户实体
type DbUser struct {
	ID       int    `gorm:"primary_key" json:"id"` //id
	Username string `json:"username"`              //用户名
	Password string `json:"password"`              //登录密码
	RoleId   int    `json:"-"`                     //角色id
	Role     DbRole `gorm:"-" json:"role"`
	Ban      int    `json:"ban"` //账号状态
}
