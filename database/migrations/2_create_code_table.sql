-- +goose Up
create table codes(
  id int(10) unsigned not null auto_increment,
  user_id int(10) unsigned not null,
  title varchar(255) not null,
  body text not null,
  status enum('public', 'private', 'limited_release'),
  created_at timestamp not null default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,

  primary key(id),
  index index_on_status(status),
  unique key uq_user_id_and_title (user_id, title),
  constraint codes_user_id_fk_user
    foreign key (user_id)
    references users(id)
    on update cascade
    on delete cascade
);

-- +goose Down
drop table codes;
