package module

type UserInfo struct {
	Id        uint32 `json:"id" gorm:"column:id"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
}

func (*UserInfo) TableName() string {
	return "t_user_info"
}

//type Register struct {
//	MobileNo string `json:"mobileNo"`
//}
