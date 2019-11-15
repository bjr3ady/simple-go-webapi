USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_item`;
CREATE TABLE `schoolboard_item` (
  `id` int(10) unique NOT NULL AUTO_INCREMENT,
  `item_id` varchar(32) unique NOT NULL COMMENT '展示项目ID(guid)',
  `name` varchar(100) unique NOT NULL COMMENT '展示项目名称',
  `link` TEXT NOT NULL COMMENT '项目链接到的目的地',
  `index` int(100) DEFAULT 0 COMMENT '项目序号',
  `sub_cate_id` varchar(32) NOT NULL COMMENT '子类别ID',
  `thumb` TEXT NULL COMMENT '项目缩略图url',
  `created_on` int(10) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='展示项目';