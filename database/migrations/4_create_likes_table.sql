-- +goose Up
create table likes(
  user_id int(10) unsigned not null,
  code_id int(10) unsigned not null,

  created_at timestamp not null default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,

  primary key (user_id, code_id),
  constraint likes_user_id_fk_users
    foreign key (user_id)
    references users(id)
    on update cascade
    on delete cascade,

  constraint likes_code_id_fk_users
    foreign key (code_id)
    references codes(id)
    on update cascade
    on delete cascade
);

-- +goose Down
drop table likes;
