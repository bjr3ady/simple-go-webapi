USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_admin`;

CREATE TABLE `schoolboard_admin` (
  `id` INT(10) unique NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(32) unique NOT NULL DEFAULT '' COMMENT '管理员ID(guid)',
  `name` varchar(100) NOT NULL unique COMMENT '管理员名称',
  `token` TEXT NOT NULL COMMENT '管理员登录Token',
  `token_expire` int(10) unsigned DEFAULT '0' COMMENT 'Token超时时间',
  `pwd` TEXT NOT NULL COMMENT '管理员密码MD5',
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(11) DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(32) DEFAULT '' COMMENT '修改人(guid)',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '管理员';