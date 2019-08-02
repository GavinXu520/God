package model

import (
	"time"
)

/******sql******
CREATE TABLE `user_link` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外建， 关联t_user_info的主键',
  `first_link_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户首要联系人姓名',
  `first_link_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '用户首要联系人手机号码',
  `second_link_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户次要联系人姓名',
  `second_link_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '用户次要联系人手机号码',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户联系人表'
******sql******/
//	用户联系人表
type UserLink struct {
	ID               int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`              //	主键Id
	UserID           int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                        //	外建， 关联t_user_info的主键
	FirstLinkName    string    `gorm:"column:first_link_name;type:varchar(32);not null" json:"first_link_name"`       //	用户首要联系人姓名
	FirstLinkMobile  string    `json:"first_link_mobile" gorm:"column:first_link_mobile;type:varchar(20);not null"`   //	用户首要联系人手机号码
	SecondLinkName   string    `gorm:"column:second_link_name;type:varchar(32);not null" json:"second_link_name"`     //	用户次要联系人姓名
	SecondLinkMobile string    `json:"second_link_mobile" gorm:"column:second_link_mobile;type:varchar(20);not null"` //	用户次要联系人手机号码
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`                   //	创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`                   //	修改时间
	Status           int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`                 //	状态
}

/******sql******
CREATE TABLE `user_login_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `login_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：password，2:手机验证码  3:第三方登陆',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '设备版本号',
  `longitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置经度',
  `latitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置纬度',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登陆历史表'
******sql******/
//	用户登陆历史表
type UserLoginHistory struct {
	ID         int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`      //
	UserID     int64     `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`       //	用户账号
	LoginType  bool      `gorm:"column:login_type;type:tinyint(1) unsigned;not null" json:"login_type"` //	0:未知，1：password，2:手机验证码  3:第三方登陆
	TerminalID int       `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`            //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode string    `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`         //	设备编码
	Version    string    `gorm:"column:version;type:varchar(10);not null" json:"version"`               //	设备版本号
	Longitude  string    `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`           //	gps 位置经度
	Latitude   string    `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`             //	gps 位置纬度
	CityID     int       `json:"city_id" gorm:"column:city_id;type:int(10) unsigned;not null"`          //	城市_id
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`           //	创建时间
}

/******sql******
CREATE TABLE `account_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外建， 关联t_user_info的主键',
  `account` varchar(30) NOT NULL DEFAULT '' COMMENT '账户',
  `nick` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `login_pwd` varchar(6) NOT NULL DEFAULT '' COMMENT '登录密码, 经过MD5加密之后的内容',
  `pay_pwd` varchar(6) NOT NULL DEFAULT '' COMMENT '支付密码, 经过MD5加密之后的内容',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `mail` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱账户',
  `img` varchar(100) NOT NULL DEFAULT '' COMMENT '头像链接',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户账户信息表'
******sql******/
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

/******sql******
CREATE TABLE `user_addr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外建， 关联t_user_info的主键',
  `country` tinyint(2) NOT NULL DEFAULT '0' COMMENT '国家',
  `city` tinyint(2) NOT NULL DEFAULT '0' COMMENT '城市',
  `home` varchar(60) NOT NULL DEFAULT '' COMMENT '家庭住址',
  `live_at` date NOT NULL DEFAULT '0000-00-00' COMMENT '居住时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户详细住址表'
******sql******/
//	用户详细住址表
type UserAddr struct {
	ID        int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"` //	主键Id
	UserID    int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           //	外建， 关联t_user_info的主键
	Country   int       `gorm:"column:country;type:tinyint(2);not null" json:"country"`           //	国家
	City      int       `gorm:"column:city;type:tinyint(2);not null" json:"city"`                 //	城市
	Home      string    `gorm:"column:home;type:varchar(60);not null" json:"home"`                //	家庭住址
	LiveAt    time.Time `gorm:"column:live_at;type:date;not null" json:"live_at"`                 //	居住时间
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`      //	创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`      //	修改时间
	Status    int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`    //	状态
}

/******sql******
CREATE TABLE `user_company` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外建， 关联t_user_info的主键',
  `company` varchar(60) NOT NULL DEFAULT '' COMMENT '就职公司',
  `trade` tinyint(2) NOT NULL DEFAULT '0' COMMENT '行业： ',
  `career` tinyint(2) NOT NULL DEFAULT '0' COMMENT '职业： ',
  `income` varchar(32) NOT NULL DEFAULT '' COMMENT '月收入',
  `phone_no` varchar(20) NOT NULL DEFAULT '' COMMENT '办公电话',
  `job_age` tinyint(4) NOT NULL DEFAULT '0' COMMENT '工龄',
  `loan_purpose` varchar(100) NOT NULL DEFAULT '' COMMENT '贷款目的',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户公司认证表'
******sql******/
//	用户公司认证表
type UserCompany struct {
	ID          int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`   //	主键Id
	UserID      int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`             //	外建， 关联t_user_info的主键
	Company     string    `gorm:"column:company;type:varchar(60);not null" json:"company"`            //	就职公司
	Trade       int       `gorm:"column:trade;type:tinyint(2);not null" json:"trade"`                 //	行业：
	Career      int       `gorm:"column:career;type:tinyint(2);not null" json:"career"`               //	职业：
	Income      string    `gorm:"column:income;type:varchar(32);not null" json:"income"`              //	月收入
	PhoneNo     string    `gorm:"column:phone_no;type:varchar(20);not null" json:"phone_no"`          //	办公电话
	JobAge      int       `gorm:"column:job_age;type:tinyint(4);not null" json:"job_age"`             //	工龄
	LoanPurpose string    `gorm:"column:loan_purpose;type:varchar(100);not null" json:"loan_purpose"` //	贷款目的
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`        //	创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`        //	修改时间
	Status      int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`      //	状态
}

/******sql******
CREATE TABLE `user_change_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `old_password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '设备版本号',
  `longitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置经度',
  `latitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置纬度',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户修改密码历史表'
******sql******/
//	用户修改密码历史表
type UserChangePasswordHistory struct {
	ID          int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`  //
	UserID      int       `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`   //	用户账号
	OldPassword string    `gorm:"column:old_password;type:varchar(50);not null" json:"old_password"` //	密码
	TerminalID  int       `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`        //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode  string    `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`     //	设备编码
	Version     string    `gorm:"column:version;type:varchar(10);not null" json:"version"`           //	设备版本号
	Longitude   string    `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`       //	gps 位置经度
	Latitude    string    `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`         //	gps 位置纬度
	CityID      int       `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`      //	城市_id
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`       //	创建时间
}

/******sql******
CREATE TABLE `user_change_pay_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `mail` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱账户',
  `old_pay_password` varchar(50) NOT NULL DEFAULT '' COMMENT '支付密码',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '设备版本号',
  `longitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置经度',
  `latitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置纬度',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户修改支付密码历史表'
******sql******/
//	用户修改支付密码历史表
type UserChangePayPasswordHistory struct {
	ID             int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`          //
	UserID         int       `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`           //	用户账号
	Mobile         string    `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`                     //	手机
	Mail           string    `gorm:"column:mail;type:varchar(32);not null" json:"mail"`                         //	邮箱账户
	OldPayPassword string    `gorm:"column:old_pay_password;type:varchar(50);not null" json:"old_pay_password"` //	支付密码
	TerminalID     int       `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`                //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode     string    `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`             //	设备编码
	Version        string    `gorm:"column:version;type:varchar(10);not null" json:"version"`                   //	设备版本号
	Longitude      string    `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`               //	gps 位置经度
	Latitude       string    `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`                 //	gps 位置纬度
	CityID         int       `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`              //	城市_id
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`               //	创建时间
}

/******sql******
CREATE TABLE `user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `first_name` varchar(12) NOT NULL DEFAULT '' COMMENT '用户真实姓',
  `last_name` varchar(12) NOT NULL DEFAULT '' COMMENT '用户真实名',
  `frontal_img` varchar(100) NOT NULL DEFAULT '' COMMENT '正面照链接',
  `hand_img` varchar(100) NOT NULL DEFAULT '' COMMENT '手拿照链接',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别, 0: 未知性别; 1: 男; 2: 女',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `marriage` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已婚, 0: no; 1: yes',
  `education` tinyint(4) NOT NULL DEFAULT '0' COMMENT '教育水平',
  `school` varchar(60) NOT NULL DEFAULT '' COMMENT '毕业学校',
  `identity_no` varchar(20) NOT NULL DEFAULT '' COMMENT '身份证件号码',
  `bank_card_no` varchar(20) NOT NULL DEFAULT '' COMMENT '银行卡号码',
  `bank_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '银行唯一标识',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户主要信息表'
******sql******/
//	用户主要信息表
type UserInfo struct {
	ID         int       `json:"-" gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`  //	主键Id
	FirstName  string    `json:"first_name" gorm:"column:first_name;type:varchar(12);not null"`     //	用户真实姓
	LastName   string    `gorm:"column:last_name;type:varchar(12);not null" json:"last_name"`       //	用户真实名
	FrontalImg string    `gorm:"column:frontal_img;type:varchar(100);not null" json:"frontal_img"`  //	正面照链接
	HandImg    string    `gorm:"column:hand_img;type:varchar(100);not null" json:"hand_img"`        //	手拿照链接
	Gender     bool      `gorm:"column:gender;type:tinyint(1);not null" json:"gender"`              //	性别, 0: 未知性别; 1: 男; 2: 女
	Birthday   time.Time `gorm:"column:birthday;type:date;not null" json:"birthday"`                //	生日
	Marriage   bool      `gorm:"column:marriage;type:tinyint(1);not null" json:"marriage"`          //	是否已婚, 0: no; 1: yes
	Education  int       `gorm:"column:education;type:tinyint(4);not null" json:"education"`        //	教育水平
	School     string    `json:"school" gorm:"column:school;type:varchar(60);not null"`             //	毕业学校
	IDentityNo string    `gorm:"column:identity_no;type:varchar(20);not null" json:"identity_no"`   //	身份证件号码
	BankCardNo string    `gorm:"column:bank_card_no;type:varchar(20);not null" json:"bank_card_no"` //	银行卡号码
	BankID     int       `gorm:"column:bank_id;type:tinyint(4);not null" json:"bank_id"`            //	银行唯一标识
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`       //	创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`       //	修改时间
	Status     int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`     //	状态
}

/******sql******
CREATE TABLE `user_register_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外建， 关联t_user_info的主键',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '设备版本号',
  `longitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置经度',
  `latitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置纬度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户注册信息表'
******sql******/
//	用户注册信息表
type UserRegisterInfo struct {
	ID         int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"` //	主键Id
	UserID     int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           //	外建， 关联t_user_info的主键
	TerminalID int       `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`       //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode string    `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`    //	设备编码
	Version    string    `gorm:"column:version;type:varchar(10);not null" json:"version"`          //	设备版本号
	Longitude  string    `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`      //	gps 位置经度
	Latitude   string    `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`        //	gps 位置纬度
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`      //	创建时间
}

/******sql******
CREATE TABLE `loan_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` varchar(10) NOT NULL DEFAULT '' COMMENT '用户名',
  `loan_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '借款的钱, 单位: 分?',
  `open_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '放款金额, 单位: 分?',
  `open_term` int(4) NOT NULL DEFAULT '0' COMMENT '放款周期',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='借款申请表'
******sql******/
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

/******sql******
CREATE TABLE `option_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr value id',
  `option_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '选项类型, 0: 未知; 1: 城市; 2: 银行名称',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `option_value` text NOT NULL COMMENT '值',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_optiuon_id` (`option_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='公共选项值表'
******sql******/
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

/******sql******
CREATE TABLE `user_bind_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `old_bank_card_no` varchar(20) NOT NULL DEFAULT '' COMMENT '旧有的银行卡号码',
  `old_bank_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '旧有的银行唯一标识',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '设备版本号',
  `longitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置经度',
  `latitude` varchar(10) NOT NULL DEFAULT '0.0' COMMENT 'gps 位置纬度',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户绑定第三方账号历史表'
******sql******/
//	用户绑定第三方账号历史表
type UserBindHistory struct {
	ID            int       `json:"-" gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`          //
	UserID        int       `json:"user_id" gorm:"column:user_id;type:bigint(20) unsigned;not null"`           //	用户账号
	OldBankCardNo string    `gorm:"column:old_bank_card_no;type:varchar(20);not null" json:"old_bank_card_no"` //	旧有的银行卡号码
	OldBankID     int       `gorm:"column:old_bank_id;type:tinyint(4);not null" json:"old_bank_id"`            //	旧有的银行唯一标识
	TerminalID    int       `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`                //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode    string    `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`             //	设备编码
	Version       string    `gorm:"column:version;type:varchar(10);not null" json:"version"`                   //	设备版本号
	Longitude     string    `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`               //	gps 位置经度
	Latitude      string    `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`                 //	gps 位置纬度
	CityID        int       `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`              //	城市_id
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`               //	创建时间
}
