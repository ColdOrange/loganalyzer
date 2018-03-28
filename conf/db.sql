DROP TABLE IF EXISTS `log`;

# The table has the same format with `format.go`.
# DO NOT EDIT.

CREATE TABLE `log` (
  `id` int(15) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(46) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  `request_method` varchar(10) DEFAULT NULL,
  `request_url` varchar(255) DEFAULT NULL,
  `http_version` varchar(10) DEFAULT NULL,
  `response_code` varchar(5) DEFAULT NULL,
  `response_time` int(10) DEFAULT NULL,
  `content_size` int(20) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `referer` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `ip` (`ip`),
  INDEX `time` (`time`),
  INDEX `request_url` (`request_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
