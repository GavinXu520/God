package module

type UserInfo struct {
	Id   uint32 `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (*UserInfo) TableName() string {
	return "t_user"
}
