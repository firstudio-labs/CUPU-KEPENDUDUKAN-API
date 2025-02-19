-- Insert dummy data into users table
INSERT INTO users (NIK, full_name, province, district, sub_district, village, full_address, coordinate, roles, username, password, created_at, updated_at)
VALUES
    ('1234567890123456', 'Ali Surya', 'Jawa Barat', 'Bandung', 'Cidadap', 'Cimahi', 'Jawa Barat, Bandung, Cidadap, Cimahi', ST_GeomFromText('POINT(107.6098 -6.9175)'), 'admin', 'aliadmin', 'password123', UNIX_TIMESTAMP(), NULL),
    ('2345678901234567', 'Budi Prasetyo', 'Jawa Timur', 'Surabaya', 'Wonokromo', 'Karang Pilang', 'Jawa Timur, Surabaya, Wonokromo, Karang Pilang', ST_GeomFromText('POINT(112.7521 -7.2997)'), 'technician', 'budi.tech', 'password456', UNIX_TIMESTAMP(), NULL),
    ('3456789012345678', 'Siti Aisyah', 'DKI Jakarta', 'Jakarta Selatan', 'Cilandak', 'Pondok Labu', 'DKI Jakarta, Jakarta Selatan, Cilandak, Pondok Labu', ST_GeomFromText('POINT(106.7893 -6.3094)'), 'region', 'siti.region', 'password789', UNIX_TIMESTAMP(), NULL),
    ('4567890123456789', 'Andi Wijaya', 'Bali', 'Denpasar', 'Denpasar Barat', 'Padangsambian Kaja', 'Bali, Denpasar, Denpasar Barat, Padangsambian Kaja', ST_GeomFromText('POINT(115.2165 -8.6693)'), 'citizens-data', 'andi.data', 'password101', UNIX_TIMESTAMP(), NULL);

-- Insert dummy data into packet_internets table
INSERT INTO packet_internets (code, name_provider, source, packet, Durations, price)
VALUES
    ('PKT001', 'Telkomsel', 'Paket Data', 'Unlimited', 2592000, 100000),
    ('PKT002', 'Indosat', 'Paket Data', '50GB', 2592000, 150000),
    ('PKT003', 'XL Axiata', 'Paket Data', '30GB', 2592000, 120000);

-- Insert dummy data into complaints table
INSERT INTO complaints (user_id, packet_internets_code, village, complaint_message, technician_note, reply, status, created_at, updated_at)
VALUES
    (1, 'PKT001', 'Cimahi', 'Internet tidak stabil', 'Perlu pengecekan lebih lanjut', 'Mohon maaf atas ketidaknyamanannya, sedang dalam proses perbaikan', 'acc', UNIX_TIMESTAMP(), NULL),
    (2, 'PKT002', 'Karang Pilang', 'Lambat dalam mengakses website', 'Sedang menunggu konfirmasi dari provider', 'Terima kasih atas laporan Anda, sedang diproses', 'acc', UNIX_TIMESTAMP(), NULL),
    (3, 'PKT003', 'Pondok Labu', 'Paket data tidak berfungsi', NULL, 'Segera diperbaiki', 'refuse', UNIX_TIMESTAMP(), NULL);

-- Insert dummy data into subs_packet table
INSERT INTO subs_packet (user_id, packet_internets_code, lifetime, payment_time, status, created_at, updated_at)
VALUES
    (2, 'PKT001', UNIX_TIMESTAMP() + 2592000, UNIX_TIMESTAMP(), 'paid', UNIX_TIMESTAMP(), NULL),
    (1, 'PKT002', UNIX_TIMESTAMP() + 2592000, UNIX_TIMESTAMP(), 'paid', UNIX_TIMESTAMP(), NULL),
    (1, 'PKT003', UNIX_TIMESTAMP() + 2592000, UNIX_TIMESTAMP(), 'unpaid', UNIX_TIMESTAMP(), NULL);



SELECT * FROM subs_packet;