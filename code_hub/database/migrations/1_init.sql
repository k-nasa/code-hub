-- +goose Up
CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  firebase_uid varchar(255) not null,
  email varchar(255),
  username varchar(255),
  icon_url varchar(255),
  created_at timestamp not null default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,

  primary key(id),
  index index_on_username(username),
  unique key firebase_uid (firebase_uid),
  unique key user_name (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
drop table users;
