USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_category`;
CREATE TABLE `schoolboard_category` (
  `id` int(10) unique NOT NULL AUTO_INCREMENT,
  `category_id` varchar(32) unique NOT NULL COMMENT '项目类别ID(guid)',
  `name` varchar(100) unique NOT NULL COMMENT '项目类别名称',
  `icon` TEXT NULL COMMENT '项目类别图标',
  `banner_bg_color` varchar(10) DEFAULT '#fff' COMMENT '项目类别栏目背景颜色',
  `thumb` TEXT NULL COMMENT '项目类别缩略图url',
  `created_on` int(10) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目类别';