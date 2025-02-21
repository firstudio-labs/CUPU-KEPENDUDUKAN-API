# DROP DATABASE  jaritmas;
CREATE DATABASE jaritmas;
USE jaritmas;


CREATE TABLE genders
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;

CREATE TABLE regions
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;

CREATE TABLE status
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;

CREATE TABLE family_status
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;

CREATE TABLE education_status
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;
CREATE TABLE jobs
(
    id   int primary key auto_increment,
    name varchar(50) not null
) ENGINE = InnoDB;

CREATE TABLE provinces
(
    id            int primary key auto_increment,
    province_code varchar(20) unique,
    name          varchar(60),
    coordinate    varchar(100)
) ENGINE = InnoDB;

CREATE TABLE districts
(
    id            int primary key auto_increment,
    district_code varchar(20) unique,
    name          varchar(60),
    coordinate    varchar(100)
) ENGINE = InnoDB;

CREATE TABLE sub_districts
(
    id                int primary key auto_increment,
    sub_district_code varchar(20) unique,
    name              varchar(60),
    coordinate        varchar(100)
) ENGINE = InnoDB;

CREATE TABLE villages
(
    id           int primary key auto_increment,
    village_code varchar(20) unique,
    name         varchar(60),
    coordinate   varchar(100)
) ENGINE = InnoDB;

CREATE TABLE Citizens
(
    id                int primary key auto_increment,
    NIK               varchar(16) unique not null,
    KK                varchar(16)        not null,
    full_name         varchar(255)       not null,
    birth_place_code  varchar(20)        not null,
    gender_id         int                not null,
    religion_id       int                not null,
    status_id         int                not null,
    family_status_id  int                not null,
    education_status  int                not null,
    job_type_id       int                not null,
    mother            varchar(255),
    father            varchar(255),
    family_head       varchar(255)       not null,
    province_code     varchar(20)        not null,
    district_code     varchar(20)        not null,
    sub_district_code varchar(20)        not null,
    village_code      varchar(20)        not null,

    FOREIGN KEY (birth_place_code) REFERENCES provinces (province_code) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (gender_id) references genders (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (religion_id) references regions (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (status_id) references status (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (family_status_id) references family_status (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (education_status) references education_status (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (job_type_id) references jobs (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (province_code) REFERENCES provinces (province_code) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (district_code) REFERENCES districts (district_code) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (sub_district_code) REFERENCES sub_districts (sub_district_code) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (village_code) REFERENCES villages (village_code) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB;
