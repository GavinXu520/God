
-- 创建数据库
CREATE DATABASE d_god;


-- 各种建表语句

use d_god;

-- 用户注册表
DROP TABLE IF EXISTS user_register_info;
create table `user_register_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  `terminal_id` int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  `devicecode` varchar(50) NOT NULL  DEFAULT '' COMMENT '设备编码',
  `version` varchar(10) NOT NULL  DEFAULT '' COMMENT '设备版本号',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `created_at`  timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at`  timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户注册信息表';



-- 用户账户信息表
DROP TABLE IF EXISTS t_account_info;
create table `t_account_info` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
	`user_id` bigint(20) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  `account` varchar(30) NOT NULL  DEFAULT '' COMMENT '账户',
  `nick` varchar(30) NOT NULL  DEFAULT '' COMMENT '昵称',
  `login_pwd` varchar(6) NOT NULL  DEFAULT '' COMMENT '登录密码, 经过MD5加密之后的内容',
  `pay_pwd` varchar(6) NOT NULL  DEFAULT '' COMMENT '支付密码, 经过MD5加密之后的内容',
  `mobile` varchar(20) NOT NULL  DEFAULT '' COMMENT '手机号码',
  `mail` varchar(32) NOT NULL  DEFAULT '' COMMENT '邮箱账户',
  `img` varchar(100) NOT NULL  DEFAULT '' COMMENT '头像链接',
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户账户信息表'; 

-- 用户主要信息表
DROP TABLE IF EXISTS t_user_info;
create table `t_user_info` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `first_name` varchar(12) NOT NULL  DEFAULT '' COMMENT '用户真实姓',
  `last_name` varchar(12) NOT NULL  DEFAULT '' COMMENT '用户真实名',
  `frontal_img` varchar(100) NOT NULL  DEFAULT '' COMMENT '正面照链接',
  `hand_img` varchar(100) NOT NULL  DEFAULT '' COMMENT '手拿照链接',
  `gender` tinyint(1) NOT NULL  DEFAULT 0 COMMENT '性别, 0: 未知性别; 1: 男; 2: 女',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `marriage` tinyint(1)  NOT NULL  DEFAULT 0 COMMENT '是否已婚, 0: no; 1: yes',
  `education` tinyint(4) NOT NULL DEFAULT '0' COMMENT '教育水平',
  `school` varchar(60) NOT NULL  DEFAULT '' COMMENT '毕业学校',
  `id_card_no` varchar(20) NOT NULL DEFAULT '' COMMENT '身份证件号码',
  `bank_card_no` varchar(20) NOT NULL DEFAULT '' COMMENT '银行卡号码',
  `bank_id` tinyint(4) NOT NULL DEFAULT 0 COMMENT '银行唯一标识',
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户主要信息表'; 



-- 用户联系人表
DROP TABLE IF EXISTS t_user_link;
create table `t_user_link` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  `first_link_name` varchar(32) NOT NULL  DEFAULT '' COMMENT '用户首要联系人姓名',
  `first_link_mobile` varchar(20) NOT NULL  DEFAULT '' COMMENT '用户首要联系人手机号码',
  `second_link_name` varchar(32) NOT NULL  DEFAULT '' COMMENT '用户次要联系人姓名',
  `second_link_mobile` varchar(20) NOT NULL  DEFAULT '' COMMENT '用户次要联系人手机号码',
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户联系人表'; 

-- 用户详细住址表
DROP TABLE IF EXISTS t_user_addr;
create table `t_user_addr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  `country` tinyint(2) NOT NULL DEFAULT 0 COMMENT '国家',
  `city` tinyint(2) NOT NULL DEFAULT 0 COMMENT '城市',
  `home` varchar(60) NOT NULL  DEFAULT '' COMMENT '家庭住址',
  `live_at` date not null default '0000-00-00' comment '居住时间',
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户详细住址表'; 


-- 用户公司认证表
DROP TABLE IF EXISTS t_user_addr;
create table `t_user_company` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` bigint(20) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  `company` varchar(60) NOT NULL  DEFAULT '' COMMENT '就职公司',
  `trade`  tinyint(2) NOT NULL  DEFAULT 0 COMMENT '行业： ',
  `career` tinyint(2) NOT NULL  DEFAULT 0 COMMENT '职业： ',
  `income` varchar(32) NOT NULL DEFAULT '' COMMENT '月收入',
  `phone_no` varchar(20) NOT NULL DEFAULT '' COMMENT '办公电话',
  `job_age` tinyint(4) NOT NULL DEFAULT 0 COMMENT '工龄',
  `loan_purpose` varchar(100) NOT NULL DEFAULT '' COMMENT '贷款目的', 
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户公司认证表'; 


-- 借款申请表
DROP TABLE IF EXISTS t_loan_info;
create table `t_loan_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `user_id` varchar(10) NOT NULL  DEFAULT '' COMMENT '用户名',
  `loan_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '借款的钱, 单位: 分?',
  `open_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '放款金额, 单位: 分?',
  `open_term`  int(4) NOT NULL DEFAULT 0 COMMENT '放款周期',
  `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='借款申请表'; 


-- 还有： 登录记录表， 枚举选项表(考虑和 城市选项合在一起)

DROP TABLE IF EXISTS `user_login_history`;
CREATE TABLE `user_login_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `login_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：password，2:手机验证码  3:第三方登陆',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `term_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `created_at`            timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登陆历史表';


DROP TABLE IF EXISTS `user_bind_history`;
CREATE TABLE `user_bind_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方接入id',
  `open_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '第三方接入类型',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `term_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `created_at`            timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户绑定第三方账号历史表';


DROP TABLE IF EXISTS `user_change_password_history`;
CREATE TABLE `user_change_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `old_password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `term_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `created_at`            timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户修改密码历史表';


DROP TABLE IF EXISTS `user_change_pay_password_history`;
CREATE TABLE `user_change_pay_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `old_pay_password` varchar(50) NOT NULL DEFAULT '' COMMENT '支付密码',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `term_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `created_at`            timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户修改支付密码历史表';


DROP TABLE IF EXISTS `attr_value_info`;
CREATE TABLE `attr_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr value id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `attr_value` text NOT NULL COMMENT '值',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_attr_id` (`attr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8  COMMENT='公共属性值表';