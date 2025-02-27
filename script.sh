#ENRTY MYSQL USE DB AND SEXUTE

SOURCE E:\Development\DEV-PROD\JARITMAS-API\db\state\indonesia_provinces.sql

SOURCE E:\Development\DEV-PROD\JARITMAS-API\db\state\ind_district.sql

SOURCE  E:\Development\DEV-PROD\JARITMAS-API\db\state\ind_sub_district.sql


docker exec -i mysql mysql -u root -p korie123 jatimas < /backup.sql


docker exec -it mysql bash
SOURCE ./docker-entrypoint-initdb.d/backup.sql
