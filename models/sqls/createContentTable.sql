USE `school_board`;
DROP TABLE IF EXISTS `schoolboard_content`;
CREATE TABLE `schoolboard_content` (
  `id` int(10) unique NOT NULL AUTO_INCREMENT,
  `content_id` varchar(32) unique NOT NULL COMMENT '项目内容ID(guid)',
  `content` TEXT NOT NULL COMMENT '项目内容',
  `sub_category_id` varchar(32) NOT NULL COMMENT '项目子类别ID(guid)',
  `video_src` TEXT NULL COMMENT '项目视频url',
  `created_on` int(10) DEFAULT NULL,
  `created_by` varchar(32) DEFAULT '' COMMENT '创建人(guid)',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目内容';