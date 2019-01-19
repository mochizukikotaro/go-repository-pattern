create table notes
(
  id serial primary key,
  content text not null,
  user_id bigint(20) unsigned,
  foreign key (user_id)
    references users(id)
);