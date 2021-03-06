package module

import "time"

//	用户账户信息表
type UserAccount struct {
	ID       int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"` //	主键Id
	Account  string `gorm:"column:account;type:varchar(30);not null" json:"account"`          //	账户
	Nick     string `gorm:"column:nick;type:varchar(30);not null" json:"nick"`                //	昵称
	LoginPwd string `gorm:"column:login_pwd;type:varchar(6);not null" json:"login_pwd"`       //	登录密码, 经过MD5加密之后的内容
	PayPwd   string `gorm:"column:pay_pwd;type:varchar(6);not null" json:"pay_pwd"`           //	支付密码, 经过MD5加密之后的内容
	Mobile   string `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`            //	手机号码
	Mail     string `gorm:"column:mail;type:varchar(32);not null" json:"mail"`                //	邮箱账户
	Img      string `gorm:"column:img;type:varchar(100);not null" json:"img"`                 //	头像链接
	Status   int    `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`    //	状态
}

func (*UserAccount) TableName() string {
	return "user_account"
}

type UserBase struct {
	ID         int       `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`  //	主键Id
	UserID     int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`            //	外建， 关联 user_account 的主键
	FirstName  string    `gorm:"column:first_name;type:varchar(12);not null" json:"first_name"`     //	用户真实姓
	LastName   string    `gorm:"column:last_name;type:varchar(12);not null" json:"last_name"`       //	用户真实名
	FrontalImg string    `gorm:"column:frontal_img;type:varchar(100);not null" json:"frontal_img"`  //	正面照链接
	HandImg    string    `gorm:"column:hand_img;type:varchar(100);not null" json:"hand_img"`        //	手拿照链接
	Gender     int       `json:"gender" gorm:"column:gender;type:tinyint(1);not null"`              //	性别, 0: 未知性别; 1: 男; 2: 女
	Birthday   time.Time `gorm:"column:birthday;type:date;not null" json:"birthday"`                //	生日
	Marriage   int       `gorm:"column:marriage;type:tinyint(1);not null" json:"marriage"`          //	是否已婚, 0: no; 1: yes
	Education  int       `gorm:"column:education;type:tinyint(4);not null" json:"education"`        //	教育水平
	School     string    `gorm:"column:school;type:varchar(60);not null" json:"school"`             //	毕业学校
	IDentityNo string    `gorm:"column:identity_no;type:varchar(20);not null" json:"identity_no"`   //	身份证件号码
	BankCardNo string    `gorm:"column:bank_card_no;type:varchar(20);not null" json:"bank_card_no"` //	银行卡号码
	BankID     int       `gorm:"column:bank_id;type:tinyint(4);not null" json:"bank_id"`            //	银行唯一标识
	Status     int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`     //	状态
}

func (*UserBase) TableName() string {
	return "user_base"
}

//	用户注册信息表
type UserRegisterInfo struct {
	ID         int    `json:"-" gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"` //	主键Id
	UserID     int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           //	外建， 关联 user_account 的主键
	TerminalID int    `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`       //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode string `json:"devicecode" gorm:"column:devicecode;type:varchar(50);not null"`    //	设备编码
	Version    string `gorm:"column:version;type:varchar(10);not null" json:"version"`          //	设备版本号
	Longitude  string `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`      //	gps 位置经度
	Latitude   string `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`        //	gps 位置纬度
}

func (*UserRegisterInfo) TableName() string {
	return "user_register_info"
}

//	用户联系人表
type UserLink struct {
	ID               int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`              //	主键Id
	UserID           int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                        //	外建， 关联 user_account 的主键
	FirstLinkName    string `gorm:"column:first_link_name;type:varchar(32);not null" json:"first_link_name"`       //	用户首要联系人姓名
	FirstLinkMobile  string `gorm:"column:first_link_mobile;type:varchar(20);not null" json:"first_link_mobile"`   //	用户首要联系人手机号码
	SecondLinkName   string `gorm:"column:second_link_name;type:varchar(32);not null" json:"second_link_name"`     //	用户次要联系人姓名
	SecondLinkMobile string `json:"second_link_mobile" gorm:"column:second_link_mobile;type:varchar(20);not null"` //	用户次要联系人手机号码
	Status           int    `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`                 //	状态
}

func (*UserLink) TableName() string {
	return "user_link"
}

//	用户登陆历史表
type UserLoginHistory struct {
	ID         int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`      //
	UserID     int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                //	外建， 关联 user_account 的主键
	LoginType  int    `gorm:"column:login_type;type:tinyint(1) unsigned;not null" json:"login_type"` //	0:未知，1：password，2:手机验证码  3:第三方登陆
	TerminalID int    `json:"terminal_id" gorm:"column:terminal_id;type:int(1);not null"`            //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode string `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`         //	设备编码
	Version    string `gorm:"column:version;type:varchar(10);not null" json:"version"`               //	设备版本号
	Longitude  string `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`           //	gps 位置经度
	Latitude   string `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`             //	gps 位置纬度
	CityID     int    `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`          //	城市_id
}

func (*UserLoginHistory) TableName() string {
	return "user_login_history"
}

//	用户详细住址表
type UserAddr struct {
	ID      int       `json:"-" gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"` //	主键Id
	UserID  int       `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`           //	外建， 关联 user_account 的主键
	Country int       `gorm:"column:country;type:tinyint(2);not null" json:"country"`           //	国家
	City    int       `gorm:"column:city;type:tinyint(2);not null" json:"city"`                 //	城市
	Home    string    `json:"home" gorm:"column:home;type:varchar(60);not null"`                //	家庭住址
	LiveAt  time.Time `gorm:"column:live_at;type:date;not null" json:"live_at"`                 //	居住时间
	Status  int       `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`    //	状态
}

func (*UserAddr) TableName() string {
	return "user_addr"
}

//	用户公司认证表
type UserCompany struct {
	ID          int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`   //	主键Id
	UserID      int    `json:"user_id" gorm:"column:user_id;type:bigint(20);not null"`             //	外建， 关联 user_account 的主键
	Company     string `gorm:"column:company;type:varchar(60);not null" json:"company"`            //	就职公司
	Trade       int    `gorm:"column:trade;type:tinyint(2);not null" json:"trade"`                 //	行业：
	Career      int    `gorm:"column:career;type:tinyint(2);not null" json:"career"`               //	职业：
	Income      string `gorm:"column:income;type:varchar(32);not null" json:"income"`              //	月收入
	PhoneNo     string `gorm:"column:phone_no;type:varchar(20);not null" json:"phone_no"`          //	办公电话
	JobAge      int    `gorm:"column:job_age;type:tinyint(4);not null" json:"job_age"`             //	工龄
	LoanPurpose string `gorm:"column:loan_purpose;type:varchar(100);not null" json:"loan_purpose"` //	贷款目的
	Status      int    `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`      //	状态
}

func (*UserCompany) TableName() string {
	return "user_company"
}

//	用户修改密码历史表
type UserChangePasswordHistory struct {
	ID          int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`  //
	UserID      int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`            //	外建， 关联 user_account 的主键
	OldPassword string `gorm:"column:old_password;type:varchar(50);not null" json:"old_password"` //	密码
	TerminalID  int    `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`        //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode  string `json:"devicecode" gorm:"column:devicecode;type:varchar(50);not null"`     //	设备编码
	Version     string `gorm:"column:version;type:varchar(10);not null" json:"version"`           //	设备版本号
	Longitude   string `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`       //	gps 位置经度
	Latitude    string `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`         //	gps 位置纬度
	CityID      int    `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`      //	城市_id
}

func (*UserChangePasswordHistory) TableName() string {
	return "user_change_password_history"
}

//	用户修改支付密码历史表
type UserChangePayPasswordHistory struct {
	ID             int    `json:"-" gorm:"primary_key;column:id;type:bigint(20) unsigned;not null"`          //
	UserID         int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                    //	外建， 关联 user_account 的主键
	Mobile         string `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`                     //	手机
	Mail           string `gorm:"column:mail;type:varchar(32);not null" json:"mail"`                         //	邮箱账户
	OldPayPassword string `gorm:"column:old_pay_password;type:varchar(50);not null" json:"old_pay_password"` //	支付密码
	TerminalID     int    `json:"terminal_id" gorm:"column:terminal_id;type:int(1);not null"`                //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode     string `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`             //	设备编码
	Version        string `gorm:"column:version;type:varchar(10);not null" json:"version"`                   //	设备版本号
	Longitude      string `json:"longitude" gorm:"column:longitude;type:varchar(10);not null"`               //	gps 位置经度
	Latitude       string `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`                 //	gps 位置纬度
	CityID         int    `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`              //	城市_id
}

func (*UserChangePayPasswordHistory) TableName() string {
	return "user_change_pay_password_history"
}

//	用户绑定第三方账号历史表
type UserBindHistory struct {
	ID            int    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`          //
	UserID        int    `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`                    //	外建， 关联 user_account 的主键
	OldBankCardNo string `gorm:"column:old_bank_card_no;type:varchar(20);not null" json:"old_bank_card_no"` //	旧有的银行卡号码
	OldBankID     int    `gorm:"column:old_bank_id;type:tinyint(4);not null" json:"old_bank_id"`            //	旧有的银行唯一标识
	TerminalID    int    `gorm:"column:terminal_id;type:int(1);not null" json:"terminal_id"`                //	平台Id, 0: 未知平台; 1: 安卓; 2: IOS;
	Devicecode    string `gorm:"column:devicecode;type:varchar(50);not null" json:"devicecode"`             //	设备编码
	Version       string `json:"version" gorm:"column:version;type:varchar(10);not null"`                   //	设备版本号
	Longitude     string `gorm:"column:longitude;type:varchar(10);not null" json:"longitude"`               //	gps 位置经度
	Latitude      string `gorm:"column:latitude;type:varchar(10);not null" json:"latitude"`                 //	gps 位置纬度
	CityID        int    `gorm:"column:city_id;type:int(10) unsigned;not null" json:"city_id"`              //	城市_id
}

func (*UserBindHistory) TableName() string {
	return "user_bind_history"
}
