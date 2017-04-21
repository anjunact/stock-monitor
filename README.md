# stock-monitor
create table products(
  id serial primary key,
  title varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  price decimal(5,2) NOT NULL
);