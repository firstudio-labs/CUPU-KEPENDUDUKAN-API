# DROP DATABASE  jaritmas;
CREATE DATABASE jaritmas;
USE jaritmas;


CREATE TABLE users
(
    id           int primary key auto_increment,
    NIK          varchar(16) unique                                    not null,
    full_name    varchar(100)                                          not null,
    province     varchar(100)                                          not null,
    district     varchar(100)                                          not null,
    sub_district varchar(100)                                          not null,
    village      varchar(100)                                          not null,
    full_address VARCHAR(255) GENERATED ALWAYS AS (
        CONCAT_WS(', ', village, sub_district, district, province)
        ) STORED, # mix from province,district, sub_district, village
    coordinate   varchar(255)                                          not null,
    roles        enum ('admin', 'technician','region','citizens-data') not null,
    username     varchar(100)                                          not null unique,
    password     varchar(100)                                          not null,
    created_at   BIGINT                                                NOT NULL COMMENT 'Unix timestamp in ms',
    updated_at   bigint,
    deleted_at   bigint                                                null default null
) ENGINE = InnoDB;

CREATE TABLE source_internets
(
    id            int primary key auto_increment,
    provider_name varchar(100) not null,
    source        varchar(100) not null
);

CREATE table packet_internets
(
    code               varchar(20) primary key unique,
    source_internet_id int    not null,
    description        varchar(255) null ,
    packet             varchar(50),
    duration           bigint not null,
    price              int    not null,
    FOREIGN KEY (source_internet_id) REFERENCES source_internets (id) on update cascade
) ENGINE = InnoDB;

#technician confuse
CREATE table complaints
(
    id                    int primary key auto_increment,
    user_id               int,
    packet_internets_code varchar(20)  not null,
    village               varchar(100) not null,
    complaint_message     text         not null,
    technician_note       varchar(100) null             default null,
    reply                 text         null,
    status                ENUM ('rejected', 'accepted') DEFAULT NULL,
    created_at            BIGINT       NOT NULL COMMENT 'Unix timestamp in ms',
    updated_at            bigint,
    deleted_at            bigint       null             default null,
    FOREIGN KEY (user_id) REFERENCES users (id) on update cascade,
    FOREIGN KEY (packet_internets_code) REFERENCES packet_internets (code) on update cascade
) ENGINE = InnoDB;

CREATE table subs_packets
(
    id                    int primary key auto_increment,
    user_id               int         not null,
    packet_internets_code varchar(20) not null,
    FOREIGN KEY (packet_internets_code) REFERENCES packet_internets (code) on update cascade,
    FOREIGN KEY (user_id) REFERENCES users (id) on update cascade,
    lifetime              BIGINT           default null, #lifetime = time.now + packetInternet.durations
    payment_time          bigint,
    status                enum ('paid', 'unpaid'),
    created_at            BIGINT      NOT NULL COMMENT 'Unix timestamp in ms',
    updated_at            bigint,
    deleted_at            bigint      null default null
) ENGINE = InnoDB;

# I don't know need or no report

select *
from users;
desc users;