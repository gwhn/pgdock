create table if not exists recipes (
	id serial primary key,
	name varchar(100) not null,
	description text,
	category_id integer references categories (id)
);
