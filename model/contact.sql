CREATE TABLE `contact` (
	company varchar(255) NOT NULL,
	focal varchar(32) NOT NULL,
	mobile varchar(20) NULL,
	startDate Date NULL,
	endDate Date NULL,
	mount float NULL,
	mail varchar(100) NULL,
	org_name varchar(100) NULL,
	id int  NOT NULL ,
	CONSTRAINT contact_pkey PRIMARY KEY (id)
);


-- fn:  FindAllContact
select * from public.contact;