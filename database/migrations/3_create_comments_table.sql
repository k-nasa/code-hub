-- +goose Up
create table comments(
  id int(10) unsigned not null auto_increment,
  user_id int(10) unsigned not null,
  code_id int(10) unsigned not null,
  body text not null,

  created_at timestamp not null default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,

  primary key(id),
  constraint comments_user_id_fk_users
    foreign key(user_id)
    references users(id)
    on update cascade
    on delete cascade,

  constraint comments_code_id_fk_users
    foreign key(code_id)
    references codes(id)
    on update cascade
    on delete cascade
);

-- +goose Down
drop table comments;
