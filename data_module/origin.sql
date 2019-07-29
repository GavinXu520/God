
-- 创建数据库
CREATE DATABASE d_god;


-- 各种建表语句

use d_god;

-- 设备信息表
DROP TABLE IF EXISTS t_device_info;
create table t_device_info (
	id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  terminal_id int(1) NOT NULL COMMENT '平台Id, 0: 未知平台; 1: 安卓; 2: IOS; ',
  devicecode varchar(50) NOT NULL  DEFAULT '' COMMENT '设备编码',
  version varchar(10) NOT NULL  DEFAULT '' COMMENT '设备版本号',
  	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


-- 登录账户表
DROP TABLE IF EXISTS t_account_info;
create table t_account_info (
	id tinyint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
	user_id tinyint(10) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  account varchar(30) NOT NULL  DEFAULT '' COMMENT '账户',
  nick varchar(30) NOT NULL  DEFAULT '' COMMENT '昵称',
  login_pwd varchar(6) NOT NULL  DEFAULT '' COMMENT '登录密码, 经过MD5加密之后的内容',
  pay_pwd varchar(6) NOT NULL  DEFAULT '' COMMENT '支付密码, 经过MD5加密之后的内容',
  mobile varchar(20) NOT NULL  DEFAULT '' COMMENT '手机号码',
  mail varchar(30) NOT NULL  DEFAULT '' COMMENT '邮箱账户',
  img varchar(50) NOT NULL  DEFAULT '' COMMENT '头像链接',
  created_at timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  updated_at timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  status tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- 用户主要信息表
DROP TABLE IF EXISTS t_user_info;
create table t_user_info (
	id tinyint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  first_name varchar(10) NOT NULL  DEFAULT '' COMMENT '用户真实姓',
  last_name varchar(10) NOT NULL  DEFAULT '' COMMENT '用户真实名',
  gender tinyint(1) NOT NULL  DEFAULT 0 COMMENT '性别, 0: 未知性别; 1: 男; 2: 女',
  birthday date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  career tinyint(2) NOT NULL  DEFAULT 0 COMMENT '职业,',
  marriage tinyint(1)  NOT NULL  DEFAULT 0 COMMENT '是否已婚, 0: no; 1: yes',
  income varchar(30) NOT NULL DEFAULT '' COMMENT '月收入',
  education tinyint(4) NOT NULL DEFAULT '0' COMMENT '教育水平',
  school varchar(60) NOT NULL  DEFAULT '' COMMENT '毕业学校',
  id_card_no varchar(20) NOT NULL DEFAULT '' COMMENT '身份证件号码',
  redit_card_no varchar(20) NOT NULL DEFAULT ''COMMENT '银行卡号码',
  created_at timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  updated_at timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  status tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- 用户详细住址
DROP TABLE IF EXISTS t_user_addr;
create table t_user_addr (
  id tinyint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  user_id tinyint(10) NOT NULL  DEFAULT 0 COMMENT '外建， 关联t_user_info的主键',
  country tinyint(1) NOT NULL DEFAULT 0 COMMENT '国家',
  city tinyint(1) NOT NULL DEFAULT 0 COMMENT '城市',
  company varchar(60) NOT NULL  DEFAULT '' COMMENT '就职公司',
  home varchar(60) NOT NULL  DEFAULT '' COMMENT '家庭住址',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- 借款申请表
DROP TABLE IF EXISTS t_loan_info;
create table t_loan_info (
  id tinyint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  name varchar(10) NOT NULL  DEFAULT '' COMMENT '用户名',
  gender tinyint(1) NOT NULL  DEFAULT 0 COMMENT '性别, 0: 未知性别; 1: 男; 2: 女',
  age tinyint(3) NOT NULL  DEFAULT 0 COMMENT '年龄',
  school varchar(60) NOT NULL  DEFAULT '' COMMENT '毕业学校',
  company varchar(60) NOT NULL  DEFAULT '' COMMENT '就职公司',
    home varchar(60) NOT NULL  DEFAULT '' COMMENT '家庭住址',
    career tinyint(2) NOT NULL  DEFAULT 0 COMMENT '职业,',
    marry tinyint(1)  NOT NULL  DEFAULT 0 COMMENT '是否已婚, 0: no; 1: yes',
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


