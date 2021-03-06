USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_func`;
CREATE TABLE `schoolboard_func` (
  `id` int(10) unique NOT NULL AUTO_INCREMENT,
  `func_id` varchar(32) unique NOT NULL COMMENT '系统功能ID(guid)',
  `name` varchar(100) unique NOT NULL COMMENT '系统功能名称',
  `created_on` int(10) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统功能';