CREATE TABLE `user` (
	employee_number varchar(255) NOT NULL,
	role_name varchar(32) NOT NULL,
	create_time varchar(20) NULL,
	update_time varchar(20) NULL,
	`password` varchar(100) NULL,
	is_active varchar(1) NULL DEFAULT 0,
	`type` varchar(2) NULL DEFAULT 0,
	mail varchar(100) NULL,
	org_id int4 NULL,
	org_name varchar(100) NULL,
	id  int NOT NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);
-- fn: FindAll
select * from user;