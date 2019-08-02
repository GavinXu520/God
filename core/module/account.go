package module

import "time"

//	用户账户信息表
type AccountInfo struct {
	ID        int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"` //	主键Id
	UserID    int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           //	外建， 关联t_user_info的主键
	Account   string    `gorm:"column:account;type:varchar(30);not null" json:"account"`          //	账户
	Nick      string    `gorm:"column:nick;type:varchar(30);not null" json:"nick"`                //	昵称
	LoginPwd  string    `gorm:"column:login_pwd;type:varchar(6);not null" json:"login_pwd"`       //	登录密码, 经过MD5加密之后的内容
	PayPwd    string    `json:"pay_pwd" gorm:"column:pay_pwd;type:varchar(6);not null"`           //	支付密码, 经过MD5加密之后的内容
	Mobile    string    `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`            //	手机号码
	Mail      string    `gorm:"column:mail;type:varchar(32);not null" json:"mail"`                //	邮箱账户
	Img       string    `gorm:"column:img;type:varchar(100);not null" json:"img"`                 //	头像链接
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`      //	创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`      //	修改时间
	Status    int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`    //	状态
}

func (*AccountInfo) TableName() string {
	return "account_info"
}
