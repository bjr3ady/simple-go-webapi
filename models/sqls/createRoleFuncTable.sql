USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_role_func`;

CREATE TABLE `schoolboard_role_func` (
  `id` INT(10) unique NOT NULL AUTO_INCREMENT,
  `role_id` varchar(32) NOT NULL DEFAULT '' COMMENT '角色ID(guid)',
  `func_id` varchar(32) NOT NULL COMMENT '系统功能ID(guid)',
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(11) unsigned DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(32) DEFAULT '' COMMENT '修改人(guid)',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '角色-系统功能';