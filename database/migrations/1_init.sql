-- +goose Up
CREATE TABLE user (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  firebase_uid VARCHAR(255) NOT NULL,
  email VARCHAR(255),
  display_name VARCHAR(255),
  photo_url VARCHAR(255),
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY `firebase_uid` (`firebase_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE user;
