DROP TABLE IF EXISTS log;

# The table has the same format with `format.go`.
# DO NOT EDIT.

CREATE TABLE log
(
  id             INT(15) UNSIGNED AUTO_INCREMENT
    PRIMARY KEY,
  ip             VARCHAR(46)  NULL,
  time           DATETIME     NULL,
  request_method VARCHAR(10)  NULL,
  request_url    VARCHAR(255) NULL,
  http_version   VARCHAR(10)  NULL,
  response_code  VARCHAR(5)   NULL,
  response_time  INT(10)      NULL,
  content_size   INT(20)      NULL,
  user_agent     VARCHAR(255) NULL,
  referer        VARCHAR(255) NULL
)
  ENGINE = InnoDB
  CHARSET = utf8;

CREATE INDEX ip
  ON log (ip);

CREATE INDEX time
  ON log (time);

CREATE INDEX request_url
  ON log (request_url);
