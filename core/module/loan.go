package module

import "time"

//	借款申请表
type LoanInfo struct {
	ID         int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"` //	主键Id
	UserID     string    `json:"user_id" gorm:"column:user_id;type:varchar(10);not null"`          //	用户名
	LoanAmount int       `gorm:"column:loan_amount;type:bigint(20);not null" json:"loan_amount"`   //	借款的钱, 单位: 分?
	OpenAmount int       `gorm:"column:open_amount;type:bigint(20);not null" json:"open_amount"`   //	放款金额, 单位: 分?
	OpenTerm   int       `gorm:"column:open_term;type:int(4);not null" json:"open_term"`           //	放款周期
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`      //	创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`      //	修改时间
	Status     int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`    //	状态
}

func (*LoanInfo) TableName() string {
	return "loan_info"
}
