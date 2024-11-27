CREATE TABLE IF NOT EXISTS task (
  id int primary key auto_increment,
  name varchar(255) not null,
  finished boolean not null default false
);
