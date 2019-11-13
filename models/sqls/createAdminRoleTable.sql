USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_admin_role`;

CREATE TABLE `schoolboard_admin_role` (
  `id` INT(10) unique NOT NULL AUTO_INCREMENT,
  `admin_id` varchar(32) NOT NULL COMMENT '管理员ID(guid)',
  `role_id` varchar(32) NOT NULL COMMENT '角色ID(guid)',
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(32) DEFAULT '' COMMENT '修改人(guid)',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '管理员-角色';