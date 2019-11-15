USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_home`;
CREATE TABLE `schoolboard_home` (
  `id` int(10) unique NOT NULL AUTO_INCREMENT,
  `home_id` varchar(32) unique NULL COMMENT '项目首页栏目ID(guid)',
  `category_id` varchar(32) NULL COMMENT '项目类别ID(guid)',
  `is_direct_link` int(1) DEFAULT 0 COMMENT '是否是直接链接项目',
  `link` TEXT NOT NULL COMMENT '链接url',
  `size_mode` int(1) DEFAULT 0 COMMENT '栏目展示方块大小',
  `index` char(10) unique COMMENT '栏目序列号',
  `created_on` int(10) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目首页栏目';