use kpt;

drop table if exists `scan_record`;

create table if not exists `scan_record` (
    `id` int unsigned auto_increment primary key,
    `uid` char(32) not null,
    `cid` char(32) not null,
    `tpt` varchar(12) not null,
    `ctime` datetime default current_timestamp 
) engine=innodb charset=utf8;

insert into `scan_record`(`uid`,`cid`,`tpt`,`ctime`) values
('1234','','36.0','2020-05-29 22:45:32'),
('1235','','36.2','2020-05-30 22:45:32'),
('1236','','36.4','2020-05-31 22:45:32'),
('1237','','36.6','2020-06-01 22:45:32'),
('1237','','36.1','2020-06-02 08:45:32');


-- drop table if exists `user`;

-- create table if not exists `user` (
--     `device_id` char(32) primary key,
--     `user_id` char(12) not null,
--     `class_id` char(8) not null,
--     `name` varchar(8) not null
-- ) engine=innodb charset=utf8mb4;

-- insert into `user`(`device_id`,`user_id`,`class_id`,`name`) values
-- ('device_id_1','2017213056','08051703','谢金锦'),
-- ('device_id_2','2017212576','08051704','王方诗');

drop table if exists `punch_record`;

create table if not exists `punch_record` (
    `id` int(11) unsigned auto_increment primary key,
    `uid` char(12) not null,
    `name` varchar(8) not null,
    `phone` varchar(12) not null, 
    `location` varchar(128) not null,
    `is_temperature_ok` boolean not null,
    `did_meet_hubei` boolean not null,
    `has_symptom` boolean not null,
    `is_family_diagnosed` boolean not null,
    `did_meet_diagnoses` boolean not null,
    `is_family_suspected` boolean not null,
    `ctime` datetime default current_timestamp 
) engine=innodb charset=utf8mb4;

insert into `punch_record`(`uid`,`name`,`phone`,`location`,`is_temperature_ok`,`did_meet_hubei`,`has_symptom`,`is_family_diagnosed`,`did_meet_diagnoses`,`is_family_suspected`) values
('2017213053','高寅','17784450780','重庆市南岸区重庆邮电大学',1,0,0,0,0,0);