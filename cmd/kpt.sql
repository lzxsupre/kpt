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

drop table if exists `rfid_record`;
create table if not exists `rfid_record` (
    `id` int unsigned auto_increment,
    `uid` char(32) not null,
    `rfid` char(32) not null,
    `type` tinyint not null,
    `ctime` datetime default current_timestamp,
    primary key(`id`)
) engine=innodb charset=utf8mb4;

drop table if exists `temp_record`;
create table if not exists `temp_record` (
    `id` int unsigned auto_increment,
    `uid` char(32) not null,
    `temp` float not null,
    `ctime` datetime default current_timestamp,
    primary key(`id`)
) engine=innodb charset=utf8mb4;

drop table if exists `user`;
create table if not exists `user` (
    `id` int unsigned auto_increment,
    `uid` char(32) not null,
    `cid` char(12) not null,
    `class_id` char(8) not null,
    `name` varchar(12) not null,
    `email` varchar(64) not null,
    `status` int(8) default 1,
    `ctime` datetime default current_timestamp,
    `mtime` datetime default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique (`cid`,`uid`,`email`)
) engine=innodb charset=utf8mb4;

insert into `user`(`uid`,`cid`,`class_id`,`name`,`email`) values
('2017213058','device_id_4','08051703','傅杰','1@xjj.pub'),
('2017213053','device_id_3','08051703','高寅','2@xjj.pub'),
('2017213056','device_id_1','08051703','谢金锦','1366723936@qq.com'),
('2017212576','device_id_2','08051704','王方诗','2898234819@qq.com');


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