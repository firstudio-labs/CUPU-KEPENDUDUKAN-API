DROP DATABASE jaritmas;
CREATE DATABASE jaritmas;
use jaritmas;

select * from indonesia_provinces;
select * from indonesia_districts;
select * from indonesia_sub_districts;
select * from indonesia_villages;


SELECT * FROM family_statuses;
SELECT * FROM ;
KEPALA KELUARGA
ISTRI
ANAK
KEPALA KELUARGA


CREATE TABLE `indonesia_provinces`  (
                                        `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                        `code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                        `meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                                        `created_at` timestamp NULL DEFAULT NULL,
                                        `updated_at` timestamp NULL DEFAULT NULL,
                                        PRIMARY KEY (`id`) USING BTREE,
                                        UNIQUE INDEX `indonesia_provinces_code_unique`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

CREATE TABLE indonesia_districts  (
                                      `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `code` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                      `province_code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                      `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                      `meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                                      `created_at` timestamp NULL DEFAULT NULL,
                                      `updated_at` timestamp NULL DEFAULT NULL,
                                      PRIMARY KEY (`id`) USING BTREE,
                                      UNIQUE INDEX `indonesia_cities_code_unique`(`code` ASC) USING BTREE,
                                      INDEX `indonesia_cities_province_code_foreign`(`province_code` ASC) USING BTREE,
                                      CONSTRAINT `indonesia_district_province_code_foreign` FOREIGN KEY (`province_code`) REFERENCES `indonesia_provinces` (`code`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 515 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;


CREATE TABLE `indonesia_sub_districts`  (
                                            `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                            `code` char(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                            `district_code` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                            `meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                                            `created_at` timestamp NULL DEFAULT NULL,
                                            `updated_at` timestamp NULL DEFAULT NULL,
                                            PRIMARY KEY (`id`) USING BTREE,
                                            UNIQUE INDEX `indonesia_districts_code_unique`(`code` ASC) USING BTREE,
                                            INDEX `indonesia_districts_city_code_foreign`(`district_code` ASC) USING BTREE,
                                            CONSTRAINT `indonesia_sub_districts_district_code_foreign` FOREIGN KEY (`district_code`) REFERENCES `indonesia_districts` (`code`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 7267 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;


CREATE TABLE `indonesia_villages`  (
                                       `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                       `code` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                       `sub_district_code` char(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                       `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                       `meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
                                       `created_at` timestamp NULL DEFAULT NULL,
                                       `updated_at` timestamp NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`) USING BTREE,
                                       UNIQUE INDEX `indonesia_villages_code_unique`(`code` ASC) USING BTREE,
                                       INDEX `indonesia_villages_district_code_foreign`(`sub_district_code` ASC) USING BTREE,
                                       CONSTRAINT `indonesia_villages_sub_district_code_foreign` FOREIGN KEY (`sub_district_code`) REFERENCES `indonesia_sub_districts` (`code`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 83810 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

select * from jobs;

select * from indonesia_districts; where name = 'TEMANGGUNG';
select * from indonesia_sub_districts where name = 'TEMANGGUNG';
select * from indonesia_sub_districts where id = 24;
SELECT *
FROM indonesia_villages
WHERE name LIKE 'CANDI R%';

SELECT *
FROM indonesia_villages
WHERE LOWER(name) LIKE LOWER('CANDIROTO%');

CANDIROTO
CANDIROTO

#distict dan sub dictict datanya ke balik

select * from indonesia_districts;
select * from indonesia_districts where name = 'CANDIROTO';

select * from indonesia_sub_districts;

SELECT * FROM indonesia_provinces ;
select * from indonesia_districts where name = 'KABUPATEN TEMANGGUNG';
select * from indonesia_sub_districts where name = 'CANDIROTO';
select * from indonesia_villages where name='MENTO';




# DROP TABLE  citizens;
select * from citizens;
3323092110770003
# DELETE FROM citizens;

DESC citizens;

select * from citizens;
select * from indonesia_districts WHERE name 'GUNUNG KIDUL';

select * from jobs;

SELECT *
FROM indonesia_districts
WHERE LOWER(name) LIKE LOWER('KABUPATEN MAS%');

SELECT count(*) FROM `citizens`

SELECT * FROM citizens WHERE kk = ;
