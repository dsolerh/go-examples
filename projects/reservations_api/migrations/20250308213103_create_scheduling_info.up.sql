create table if not exists scheduling_info (
	table_id serial primary key,
	capacity smallint not null
);