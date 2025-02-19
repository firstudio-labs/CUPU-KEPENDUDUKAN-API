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
    full_address varchar(255)                                          not null, # mix from province,district, sub_district, village
    coordinate   point                                                 not null,
    roles        enum ('admin', 'technician','region','citizens-data') not null,
    username     varchar(100)                                          not null unique,
    password     varchar(100)                                          not null,
    created_at   bigint                                                not null not null,
    updated_at   bigint,
    deleted_at   bigint                                                null default null
) ENGINE = InnoDB;

CREATE table packet_internets
(
    code          varchar(20) primary key unique,
    name_provider varchar(100) not null,
    source        varchar(100),
    packet        varchar(50),
    Durations     bigint       not null,
    price         int          not null
) ENGINE = InnoDB;

#technician confuse
CREATE table complaints
(
    id                    int primary key auto_increment,
    user_id               int,
    packet_internets_code varchar(20)  not null,
    village               varchar(100) not null,
    complaint_message     text         not null,
    technician_note       varchar(100) null default null,
    reply                 text         null,
    status                enum ('refuse', 'acc'),
    created_at            bigint       not null,
    updated_at            bigint,
    deleted_at            bigint       null default null,
    FOREIGN KEY (user_id) REFERENCES users (id) on update cascade,
    FOREIGN KEY (packet_internets_code) REFERENCES packet_internets (code) on update cascade
) ENGINE = InnoDB;

CREATE table subs_packet
(
    id                    int primary key auto_increment,
    user_id               int         not null,
    packet_internets_code varchar(20) not null,
    FOREIGN KEY (packet_internets_code) REFERENCES packet_internets (code) on update cascade,
    FOREIGN KEY (user_id) REFERENCES users (id) on update cascade,
    lifetime              BIGINT           default null, #lifetime = time.now + packetInternet.durations
    payment_time          bigint,
    status                enum ('paid', 'unpaid'),
    created_at            bigint      not null,
    updated_at            bigint,
    deleted_at            bigint      null default null
) ENGINE = InnoDB;

# I don't know need or no report