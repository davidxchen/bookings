create table users (
	id serial primary key,
	first_name varchar(255) not null,
	last_name varchar(255) not null,
	email varchar(255) not null,
	password varchar(255) not null,
	create_at timestamp not null,
	update_at timestamp not null,
	access_level int not null
);

create table reservations (
	id serial primary key,
	first_name varchar(255) null,
	last_name varchar(255) not null,
	email varchar(255) not null,
	phone varchar(255) null,
	start_date date not null,
	end_date date not null,
	room_id not null,
	create_at timestamp not null,
	update_at timestamp not null,
	foreign key (room_id)
		references rooms (id)
);

create table rooms (
	id serial primary key,
	room_name varchar(255) not null,
	create_at timestamp not null,
	update_at timestamp not null
);

create table restrictions (
	id serial primary key,
	restriction_name varchar(255) not null,
	create_at timestamp not null,
	update_at timestamp not null
);

create table room_restrictions (
	id serial primary key,
	start_date date not null,
	end_date date not null,
	room_id int not null,
	reservation_id int not null,
	create_at timestamp not null,
	update_at timestamp not null,
	restriction_id int not null,
	foreign key (room_id)
		references rooms (id),
	foreign key (reservation_id)
		references reservations (id),
	foreign key (restriction_id)
		references restrictions (id)
);