
-- content
-- part 1, pay center
-- part 2, virtual product/delivery center
-- part 3, user center
-- part 4, product attr/option/category

-- part 1, pay center

-- 支付信息表， 在第一期已经实现， 目前由订单中心维护

alter table `pay` add column `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id' after `id`;
alter table `pay` add column `pay_status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态, 0:未支付, 1:已成功支付, 2:支付失败' after `pay_seq`; 
alter table `pay` MODIFY column `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信..., 参考pay_type 表';
alter table `pay` MODIFY column `pay_account` varchar(100) NOT NULL DEFAULT '' COMMENT '支付账号, 第三方的openid';
alter table `pay` change column `order_id` `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id';
alter table `pay` add column `notify_status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态, 0:未通知, 1:已通知' after `pay_status`; 

alter table `order_info` add column `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信..., 参考pay_type 表' after `pay_status`;
alter table `order_info` MODIFY column `order_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态0:等待支付，1：支付成功（兼容保留不再使用），2：用户取消，3：订单失效，4：已发货，5：已退费， 6：发货中 ,7: 支付中，8:等待发货,9:发货失败,10:退费中,13:交易完成';

alter table `order_sku_epay` MODIFY column `epay_id` varchar(255) NOT NULL DEFAULT '' COMMENT '优惠_id, 商品的（非现金）券_id';

DROP TABLE IF EXISTS `pay_action`;
CREATE TABLE `pay_action` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
`pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id', 
`pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
`pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额单元分',
`pay_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
`action` int(11) NOT NULL DEFAULT '0' COMMENT '0:None, 1:支付, 2:关闭支付 99:其他', 
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order_id` (`pay_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='支付记录表';
/*
DROP TABLE IF EXISTS `pay_order_info`;
CREATE TABLE `pay_order_info` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id', 
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id', 
`pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额单元分',
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order_id` (`order_id`),
KEY `idx_pay_order_id` (`pay_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='支付订单关系表';
*/

DROP TABLE IF EXISTS `refund`;
CREATE TABLE `refund` (
`id` bigint(11) NOT NULL AUTO_INCREMENT,
`app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单 id',
`pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
`refund_id` varchar(32) NOT NULL DEFAULT '' COMMENT '退费 id', 
`refund_seq` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方退费流水 id', 
`refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退费金额',
`refund_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '退费时间', 
`refund_status` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0:未退费, 1:已成功退费, 2:退费失败', 
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态', 
PRIMARY KEY (`Id`),
KEY `idx_order_id` (`order_id`),
KEY `idx_refund_id` (`refund_id`),
KEY `idx_refund_seq` (`refund_seq`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='退费信息表';

DROP TABLE IF EXISTS `refund_action`;
CREATE TABLE `refund_action` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id', 
`pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
`refund_id` varchar(32) NOT NULL DEFAULT '' COMMENT '退费 id', 
`refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退费金额',
`refund_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '退费时间',
`action` int(11) NOT NULL DEFAULT '0' COMMENT '0:退费 99:其他', 
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order_id` (`order_id`),
KEY `idx_refund_id` (`refund_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='退费记录表';


DROP TABLE IF EXISTS `pay_type`;
CREATE TABLE `pay_type` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`type` int(11) NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信',
`name` varchar(50) DEFAULT '' COMMENT '支付名称',
`remark` varchar(255) NOT NULL DEFAULT '',
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态', 
PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='支付类型表';

insert into `pay_type` (`id`,`type`,`name`,`remark`) values (1, 1,  '支付宝WAP', 'WAP发起支付'); 
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (2, 10, '支付宝生活号', '生活号应用');
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (3, 11, '支付宝', 'APP SDK发起支付');
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (4, 12, '支付宝扫码', '扫码发起支付');

insert into `pay_type` (`id`,`type`,`name`,`remark`) values (5, 2,  '微信WAP', 'WAP发起支付');
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (6, 20, '微信公众号', '公众号应用');
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (7, 21, '微信', 'APP SDK发起支付');
insert into `pay_type` (`id`,`type`,`name`,`remark`) values (8, 22, '微信扫码', '扫码发起支付');

DROP TABLE IF EXISTS `system_param`;
CREATE TABLE `system_param` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(32) NOT NULL COMMENT '参数名字',
  `value` varchar(255) NOT NULL COMMENT '参数值',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统静态、动态参数表';

insert into `system_param`  (`id`,`name`,`value`) values (1, 'lastPollPayResultTime', '1494493167');  
   
-- part 2, virtual product/delivery center

DROP TABLE IF EXISTS `product_exchange_code`;
CREATE TABLE `product_exchange_code` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id', 
`sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id', 
`exchange_code` varchar(64) NOT NULL DEFAULT '' COMMENT '兑换码',
`flag` int NOT NULL DEFAULT '0' COMMENT '0:未使用，1：已使用',
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order_id` (`order_id`),
KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='订单商品兑换码表';


alter table product_info drop column is_member;
alter table product_info add column delivery_type int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发货类型' after `score`;
alter table product_info add column auto_delivery tinyint(1) NOT NULL DEFAULT '0'  COMMENT '是否自动发货0：不自动发货 1：自动发货' after `score`;
alter table product_info add column virtual_product tinyint(1) NOT NULL DEFAULT '0'  COMMENT '是否虚拟商品0：普通商品 1：虚拟商品' after `score`;

alter table order_info add column delivery_type int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发货类型' after `pay_status`;
alter table order_info add column auto_delivery tinyint(1) NOT NULL DEFAULT '0'  COMMENT '是否自动发货0：不自动发货 1：自动发货' after `pay_status`;
alter table order_info add column virtual_product tinyint(1) NOT NULL DEFAULT '0'  COMMENT '是否虚拟商品0：普通商品 1：虚拟商品' after `pay_status`;
alter table order_info add column `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单id' after `order_id`;
alter table order_info add column channel_id int NOT NULL DEFAULT '0'  COMMENT '商品购买渠道' after `pay_status`;
alter table order_info add column is_from_cart tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否来源购物车' after `pay_status`;
alter table order_info add column `delivery_time` timestamp not null default '0000-00-00 00:00:00' comment '发货时间' after `pay_status`;

-- 商品信息表中，新增一个商品渠道字段，目前区别是否是会员商品, product_attr中的会员属性attr_id=3废弃
alter table product_info add column  `product_source` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品渠道信息 0：普通商品 1：会员商品' after `score`;
-- 数据移植， 更新product_info 会员商品标志
update product_info p, product_attr a set product_source=1 where p.id=a.product_id and a.attr_id=3;


DROP TABLE IF EXISTS `order_product_ext_info`;
CREATE TABLE `order_product_ext_info` (
`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id', 
`product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id', 
`sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
`product_ext_info` text NOT NULL COMMENT '商扩展信息json串，发货附加信息等由业务自行设置和解析',
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order_product_sku` (`order_id`,`product_id`,`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单商品扩展信息';


DROP TABLE IF EXISTS `order_delivery_result_info`;
CREATE TABLE `order_delivery_result_info` (
`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
`order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id', 
`delivery_result_info` text NOT NULL COMMENT '发货结果信息json串，由业务自行设置和解析',
`created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
`updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
PRIMARY KEY (`id`),
KEY `idx_order` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单发货结果信息';

/*
DROP TABLE IF EXISTS `delivery_type`;
CREATE TABLE `delivery_type` (
  `id` bigint(10) unsigned NOT NULL COMMENT 'delivery type id', --  AUTO_INCREMENT, it begins with 0
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='发货类型定义表';

insert into `delivery_type` values(0,'线下发货','无特定发货动作，例如线下发货',now(),now(),0);
insert into `delivery_type` values(1,'自定义码','自定义码',now(),now(),0);
insert into `delivery_type` values(2,'采购码','采购码',now(),now(),0);
insert into `delivery_type` values(3,'电影','电影',now(),now(),0);
insert into `delivery_type` values(4,'蜂助手油卡','蜂助手油卡',now(),now(),0);
insert into `delivery_type` values(5,'蜂助手话费','蜂助手话费',now(),now(),0);
insert into `delivery_type` values(6,'蜂助手流量','	蜂助手流量',now(),now(),0);
insert into `delivery_type` values(7,'万里通','万里通',now(),now(),0);


DROP TABLE IF EXISTS `product_source_info`;
CREATE TABLE `product_source_info` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'product_source id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品渠道定义表';

insert into `product_source_info` values(1,'普通商品','普通商品',now(),now(),0);

DROP TABLE IF EXISTS `channel_id_info`;
CREATE TABLE `channel_id_info` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'channel id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='销售渠道定义表';

insert into `channel_id_info` values(1000,'ios','ios',now(),now(),0);
insert into `channel_id_info` values(1001,'android','android',now(),now(),0);
insert into `channel_id_info` values(1002,'微信公众号','微信公众号',now(),now(),0);
*/
-- part 3, user center


drop table if exists user_info;


-- ----------------------------
-- Table structure for `user_info`
-- ----------------------------

CREATE TABLE `user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户账号',
  `account_name` varchar(30) NOT NULL DEFAULT '' COMMENT '账号名，只支持英文数字',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `nick_name` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `pay_password` varchar(50) NOT NULL DEFAULT '' COMMENT '支付密码',
  `head_icon` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `mail` varchar(50) NOT NULL DEFAULT '‘’' COMMENT '注册邮箱',
  `grade_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '等级',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `real_name` varchar(30) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `marriage` tinyint(1) NOT NULL DEFAULT '0' COMMENT '婚姻状况',
  `income` varchar(30) NOT NULL DEFAULT '' COMMENT '月收入',
  `identity` varchar(50) NOT NULL DEFAULT '' COMMENT '身份证',
  `education` tinyint(4) NOT NULL DEFAULT '0' COMMENT '教育',
  `career` smallint(6) NOT NULL DEFAULT '0' COMMENT '行业',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';

DROP TABLE IF EXISTS `user_register_info`;
create table `user_register_info`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '注册ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `term_type` tinyint unsigned NOT NULL DEFAULT '0'  COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float  not null default '0.0' COMMENT 'gps 位置经度',
  `latitude` float  not null default '0.0' COMMENT 'gps 位置纬度',
  `carrier` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '运营商',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户注册信息表';


DROP TABLE IF EXISTS `user_open_id_info`;
CREATE TABLE `user_open_id_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT ,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方接入id',
  `open_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '第三方接入类型',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户第三方账号绑定信息表';

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


-- part 4, product attr/option

alter table attr_info add column  `attr_style` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '属性的值输入风格 0:文本输入, 1：单选, 2：多选' after `remark`;

alter table attr_info drop column attr_set_id;
drop table attr_set_info;

delete from attr_info;

INSERT INTO `attr_info` VALUES (1, '商品详情', '商品详情', 0,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (2, '购买须知', '购买须知', 0,  now(), now(), 0);
-- 3 isMember is discarded
INSERT INTO `attr_info` VALUES (4, '显示销量倍数', '显示销量倍数', 0,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (5, '短信模板ID', '短信模板ID', 0,  now(), now(), 0);


INSERT INTO `attr_info` VALUES (6, 'productType', '商品(订单)显示类型', 1,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (7, 'cardCategory', '商品适用卡品类', 1,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (8, '手机', '用户自定义发货字段手机输入项', 0,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (9, '油卡卡号', '自定义发货字段油卡卡号输入项', 0,  now(), now(), 0);
INSERT INTO `attr_info` VALUES (10, '姓名', '自定义发货字段姓名输入项', 0, now(), now(), 0);
INSERT INTO `attr_info` VALUES (11, '身份证', '自定义发货字段身份证输入项', 0, now(), now(), 0);
INSERT INTO `attr_info` VALUES (12, '发货地址', '用户自定义发货字段地址输入项', 0, now(), now(), 0);
INSERT INTO `attr_info` VALUES (13, '发货备注', '用户自定义发货字段其他输入项', 0, now(), now(), 0);
INSERT INTO `attr_info` VALUES (14, '是否支持退货', '是否支持退货', 1, now(), now(), 0);
INSERT INTO `attr_info` VALUES (15, '是否支持退款', '是否支持退款', 1, now(), now(), 0);
INSERT INTO `attr_info` VALUES (16, '品牌', '品牌', 1, now(), now(), 0);

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

insert into `attr_value_info` values(1,6,'1',now(),now(),0);
insert into `attr_value_info` values(2,6,'2',now(),now(),0);
insert into `attr_value_info` values(3,6,'3',now(),now(),0);
insert into `attr_value_info` values(4,6,'4',now(),now(),0);
insert into `attr_value_info` values(5,6,'5',now(),now(),0);
insert into `attr_value_info` values(6,7,'1',now(),now(),0);
insert into `attr_value_info` values(7,7,'2',now(),now(),0);

DROP TABLE IF EXISTS `attr_type`;
CREATE TABLE `attr_type` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr type id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='属性类型定义表';


insert into `attr_type` values(1,'通用复杂属性','比如商品详情，购买须知',now(),now(),0); 
insert into `attr_type` values(2,'系统后台属性','比如显示销量倍数',now(),now(),0); 
insert into `attr_type` values(3,'商品列表属性','在商品列表使用的属性',now(),now(),0);
insert into `attr_type` values(4,'订单列表属性','在订单列表使用的属性，比如是否有商品扩展信息，是否有扩展发货结果信息',now(),now(),0); 
insert into `attr_type` values(5,'下单用户输入属性','比如充值类商品需要提供用户手机号，卡号',now(),now(),0);

insert into `attr_type` values(7,'服务条款','商品政策条款',now(),now(),0);
insert into `attr_type` values(8,'商品参数描述','商品参数描述',now(),now(),0);
insert into `attr_type` values(9,'F&Q','商品F&Q',now(),now(),0);

insert into `attr_type` values(10,'内部资源消费属性','内部资源消费属性',now(),now(),0);

DROP TABLE IF EXISTS `category_attr_info`;
CREATE TABLE `category_attr_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `attr_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `category_id` int(10) unsigned NOT NULL  COMMENT '分类id',
  `attr_type_id` int(10) unsigned NOT NULL  COMMENT '属性类型id',
  `attr_mandatory` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '必须输入: 0:非必须, 1：必须',

  `created_at`    timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`    timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='分类属性定义表';



DROP TABLE IF EXISTS `category_attr_value_info`;
CREATE TABLE `category_attr_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr value id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `category_id` int(10) unsigned NOT NULL  COMMENT '分类id',
  `attr_value` text NOT NULL COMMENT '值',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类属性值表';

DROP TABLE IF EXISTS `product_attr_type`;
CREATE TABLE `product_attr_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attr_type_id` int(10) unsigned NOT NULL  COMMENT '属性类型id',
  `created_at`    timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`    timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='商品属性类型表';



DROP TABLE IF EXISTS `order_product_attr`;
CREATE TABLE `order_product_attr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `attr_type_id` int(10) unsigned NOT NULL  COMMENT '属性类型id',
  `attr_value` text NOT NULL COMMENT '属性值',
  `attr_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_product_type_id` (`order_id`,`product_id`,`attr_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单商品属性值表';

alter table option_info drop column option_set_id;
drop table option_set_info;
DROP TABLE IF EXISTS `category_option_info`;
CREATE TABLE `category_option_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '选项 id',
  `option_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `category_id` int(10) unsigned NOT NULL  COMMENT '分类id',
  `created_at`    timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`    timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='分类选项定义表';

DROP TABLE IF EXISTS `category_option_value_info`;
CREATE TABLE `category_option_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '选项value id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '选项 id',
  `category_id` int(10) unsigned NOT NULL  COMMENT '分类id',
  `option_value` text NOT NULL COMMENT '值',
  `created_at`            timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at`            timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类选项值表';

DROP TABLE IF EXISTS `front_category_info`;
CREATE TABLE `front_category_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `category_name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '缩略图',
  `sort_order` int(3) NOT NULL DEFAULT '0',
  `created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='前端分类表';

DROP TABLE IF EXISTS `front_category_mapping`;
CREATE TABLE `front_category_mapping` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `front_category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前端分类id',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前端分类类型: 0:关联后端分类, 1:关联活动id',
  `related_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关联对象id, 例如：后端分类id', 
  `created_at` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='前端分类表映射表';

-- need to do data migration for score_mall
alter table `category_info` drop column `image`;
alter table `category_info` drop column `sort_order`;


alter table `shopping_cart` add column `sku_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku名字' after `sku_id`;
alter table `shopping_cart` add column `image` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku缩略图' after `sku_id`;
alter table `shopping_cart` add column `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名字' after `sku_id`;
