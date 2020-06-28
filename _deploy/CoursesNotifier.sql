-- DDL for CoursesNotifier
-- Machine Generated

CREATE DATABASE IF NOT EXISTS coursenotifier;

USE coursenotifier;

CREATE TABLE IF NOT EXISTS course
(
    cid      char(32)    not null
        primary key,
    name     varchar(32) null,
    teacher  varchar(20) null,
    location varchar(50) null,
    begin    char(5)     null,
    end      char(5)     null,
    week     varchar(20) null,
    time     char(5)     null
) charset = utf8;


CREATE TABLE IF NOT EXISTS coursetaking
(
    sid char(12) not null,
    cid char(32) not null,
    primary key (sid, cid)
);

CREATE TABLE IF NOT EXISTS student
(
    sid        char(12) not null
        primary key,
    pwd        char(16) null,
    wxuser     char(64) null,
    createtime bigint   null,
    constraint student_wxuser_uindex
        unique (wxuser)
);


CREATE TABLE IF NOT EXISTS current
(
    termbegin date default '2020-02-17' not null
);


alter table student
    add examresults nvarchar(1023) default '' null;

-- INSERT INTO current (termbegin) VALUES ('2020-02-17');