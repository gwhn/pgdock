create table if not exists categories (
	id serial primary key,
	name varchar(100) unique not null
);
