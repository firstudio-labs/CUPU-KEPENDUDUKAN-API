SELECT *
FROM Citizens;

INSERT INTO genders (name) VALUES
                              ('Laki-laki'),
                              ('Perempuan');

INSERT INTO regions (name) VALUES
                              ('Islam'),
                              ('Kristen'),
                              ('Hindu'),
                              ('Buddha'),
                              ('Konghucu');

INSERT INTO status (name) VALUES
                              ('Belum Kawin'),
                              ('Kawin'),
                              ('Janda'),
                              ('Duda');

INSERT INTO family_status (name) VALUES
                                     ('Kepala Keluarga'),
                                     ('Anggota Keluarga');

INSERT INTO education_status (name) VALUES
                                        ('Tidak Sekolah'),
                                        ('SD'),
                                        ('SMP'),
                                        ('SMA'),
                                        ('D3'),
                                        ('S1'),
                                        ('S2'),
                                        ('S3');

INSERT INTO jobs (name) VALUES
                            ('Petani'),
                            ('Pedagang'),
                            ('Pegawai Negeri'),
                            ('Wirausaha'),
                            ('Buruh'),
                            ('Dokter'),
                            ('Guru'),
                            ('Pengacara');

INSERT INTO provinces (province_code, name, coordinate) VALUES
                                                           ('P01', 'Jawa Barat', '06.5931° S, 106.7895° E'),
                                                           ('P02', 'Jawa Tengah', '07.0048° S, 110.3770° E'),
                                                           ('P03', 'Yogyakarta', '07.8046° S, 110.3652° E');

INSERT INTO districts (district_code, name, coordinate) VALUES
                                                           ('D01', 'Bandung', '06.9175° S, 107.6191° E'),
                                                           ('D02', 'Semarang', '06.9664° S, 110.4202° E'),
                                                           ('D03', 'Yogyakarta', '07.8017° S, 110.3644° E');

INSERT INTO sub_districts (sub_district_code, name, coordinate) VALUES
                                                                   ('SD01', 'Cicendo', '06.9175° S, 107.6191° E'),
                                                                   ('SD02', 'Gajahmungkur', '06.9664° S, 110.4202° E'),
                                                                   ('SD03', 'Kotagede', '07.8017° S, 110.3644° E');

INSERT INTO villages (village_code, name, coordinate) VALUES
                                                         ('V01', 'Cigugur', '06.9300° S, 107.6000° E'),
                                                         ('V02', 'Tlogomas', '06.9600° S, 110.4100° E'),
                                                         ('V03', 'Kraton', '07.8050° S, 110.3645° E');

INSERT INTO Citizens (NIK, KK, full_name, birth_place_code, gender_id, religion_id, status_id, family_status_id, education_status, job_type_id, mother, father, family_head, province_code, district_code, sub_district_code, village_code)
VALUES
    ('3201042345678904', '3201042345678904', 'Yuni Kurniawati', 'P01', 2, 1, 1, 1, 7, 2, 'Siti Mulyani', 'Mukti', 'Yuni Kurniawati', 'P01', 'D01', 'SD01', 'V01'),
    ('3201052345678905', '3201052345678905', 'Rizky Anwar', 'P02', 1, 2, 2, 1, 5, 4, 'Elisa', 'Hamid', 'Rizky Anwar', 'P02', 'D02', 'SD02', 'V02'),
    ('3201062345678906', '3201062345678906', 'Nina Oktaviani', 'P03', 2, 3, 3, 2, 4, 7, 'Lili', 'Joko', 'Nina Oktaviani', 'P03', 'D03', 'SD03', 'V03'),
    ('3201072345678907', '3201072345678907', 'Diana Putri', 'P01', 2, 4, 1, 2, 6, 3, 'Fika', 'Taufik', 'Diana Putri', 'P01', 'D01', 'SD01', 'V01'),
    ('3201082345678908', '3201082345678908', 'Siti Aminah', 'P02', 2, 5, 2, 2, 3, 5, 'Alia', 'Abdurrahman', 'Siti Aminah', 'P02', 'D02', 'SD02', 'V02'),
    ('3201092345678909', '3201092345678909', 'Sabrina Cahya', 'P03', 2, 1, 3, 1, 8, 2, 'Erni', 'Imam', 'Sabrina Cahya', 'P03', 'D03', 'SD03', 'V03'),
    ('3201102345678910', '3201102345678910', 'Iwan Setiawan', 'P01', 1, 2, 1, 1, 4, 6, 'Nina', 'Mukti', 'Iwan Setiawan', 'P01', 'D01', 'SD01', 'V01'),
    ('3201112345678911', '3201112345678911', 'Lina Lestari', 'P02', 2, 3, 2, 2, 6, 1, 'Sari', 'Budi', 'Lina Lestari', 'P02', 'D02', 'SD02', 'V02'),
    ('3201122345678912', '3201122345678912', 'Fajar Dwi', 'P03', 1, 4, 3, 1, 7, 3, 'Ratna', 'Suryanto', 'Fajar Dwi', 'P03', 'D03', 'SD03', 'V03'),
    ('3201132345678913', '3201132345678913', 'Ariani Putri', 'P01', 2, 5, 1, 2, 5, 4, 'Arini', 'Bowo', 'Ariani Putri', 'P01', 'D01', 'SD01', 'V01'),
    ('3201142345678914', '3201142345678914', 'Dedi Saputra', 'P02', 1, 3, 2, 1, 3, 5, 'Siti', 'Teddy', 'Dedi Saputra', 'P02', 'D02', 'SD02', 'V02'),
    ('3201152345678915', '3201152345678915', 'Taufik Hidayat', 'P03', 1, 2, 3, 1, 6, 3, 'Sari', 'Hendra', 'Taufik Hidayat', 'P03', 'D03', 'SD03', 'V03'),
    ('3201162345678916', '3201162345678916', 'Budi Santoso', 'P01', 1, 1, 1, 1, 6, 4, 'Mita', 'Suryadi', 'Budi Santoso', 'P01', 'D01', 'SD01', 'V01'),
    ('3201172345678917', '3201172345678917', 'Dewi Lestari', 'P02', 2, 3, 2, 2, 4, 5, 'Ria', 'Toto', 'Dewi Lestari', 'P02', 'D02', 'SD02', 'V02'),
    ('3201182345678918', '3201182345678918', 'Toni Wijaya', 'P03', 1, 4, 3, 1, 3, 2, 'Siti', 'Alim', 'Toni Wijaya', 'P03', 'D03', 'SD03', 'V03');


SELECT * FROM Citizens;
DELETE FROM citizens WHERE NIK = '3201182345678918';
