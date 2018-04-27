# The table has the same format with constants in `format.go`.
# DO NOT EDIT.

DROP TABLE IF EXISTS log;

CREATE TABLE `log`
(
  `id`             int(15) unsigned NOT NULL AUTO_INCREMENT
    PRIMARY KEY,
  `ip`             varchar(46)      DEFAULT NULL,
  `time`           datetime         DEFAULT NULL,
  `request_method` varchar(10)      DEFAULT NULL,
  `url_path`       varchar(255)     DEFAULT NULL,
  `url_query`      varchar(255)     DEFAULT NULL,
  `url_is_static`  tinyint(1)       DEFAULT '0',
  `http_version`   varchar(10)      DEFAULT NULL,
  `response_code`  int(3) unsigned  DEFAULT NULL,
  `response_time`  int(15) unsigned DEFAULT NULL,
  `content_size`   int(15) unsigned DEFAULT NULL,
  `ua_browser`     varchar(255)     DEFAULT NULL,
  `ua_os`          varchar(255)     DEFAULT NULL,
  `ua_device`      varchar(255)     DEFAULT NULL,
  `referer_site`   varchar(255)     DEFAULT NULL,
  `referer_path`   varchar(255)     DEFAULT NULL,
  `referer_query`  varchar(255)     DEFAULT NULL,

  INDEX `ip` (`ip`),
  INDEX `time` (`time`),
  INDEX `url_path` (`url_path`),
  INDEX `url_is_static` (`url_is_static`)
)
  ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
