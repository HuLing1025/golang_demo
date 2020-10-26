package bean

// 角色实体
type DbRole struct {
	ID         int    `gorm:"primary_key" json:"id"`               //角色id
	Role       string `gorm:"column:role" json:"role"`             //角色名称
	Permission string `gorm:"column:permission" json:"permission"` //角色权限
}
