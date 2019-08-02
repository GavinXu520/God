package module

import "time"

//	公共选项值表
type OptionValueInfo struct {
	ID          int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`          //	attr value id
	OptionType  bool      `gorm:"column:option_type;type:tinyint(1);not null" json:"option_type"`            //	选项类型, 0: 未知; 1: 城市; 2: 银行名称
	OptionID    int       `gorm:"index;column:option_id;type:bigint(20) unsigned;not null" json:"option_id"` //	attr id
	OptionValue string    `gorm:"column:option_value;type:text;not null" json:"option_value"`                //	值
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`               //	创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`               //	修改时间
	Status      int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`             //	状态
}

func (*OptionValueInfo) TableName() string {
	return "option_value_info"
}
