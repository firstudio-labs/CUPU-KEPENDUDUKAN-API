INSERT INTO users (NIK, full_name, province, district, sub_district, village, coordinate, roles, username, password, created_at)
VALUES
    ('1234567890123456', 'John Doe', 'Central Java', 'Semarang', 'Tembalang', 'Kaliwuluh', '107.2028, -6.9759', 'admin', 'john.doe', 'hashedpassword123', 1672531199000),
    ('2345678901234567', 'Jane Smith', 'Yogyakarta', 'Sleman', 'Ngaglik', 'Sendangadi', '110.4077, -7.7107', 'technician', 'jane.smith', 'hashedpassword456', 1672531200000),
    ('3456789012345678', 'Ali Alamsyah', 'Bali', 'Denpasar', 'Denpasar Barat', 'Pemecutan Kelod', '115.2210, -8.6570', 'region', 'ali.alamsyah', 'hashedpassword789', 1672531201000),
    ('4567890123456789', 'Rina Hartati', 'West Java', 'Bandung', 'Cicendo', 'Cipaganti', '107.6191, -6.9175', 'citizens-data', 'rina.hartati', 'hashedpassword101', 1672531202000),
    ('5678901234567890', 'Budi Santoso', 'East Java', 'Surabaya', 'Gubeng', 'Ketintang', '112.7521, -7.2689', 'technician', 'budi.santoso', 'hashedpassword102', 1672531203000);

INSERT INTO source_internets (provider_name, source)
VALUES
    ('Telkom', 'Fiber Optic'),
    ('Indosat', '4G LTE'),
    ('XL Axiata', 'LTE-A'),
    ('Smartfren', '5G'),
    ('First Media', 'Fiber Optic');

INSERT INTO packet_internets (code, source_internet_id, description, packet, duration, price)
VALUES
    ('P001', 1, 'Fast internet with 100Mbps download speed', '100Mbps', 2592000, 500000),
    ('P002', 2, 'Affordable 4G internet package', '4G Unlimited', 8640000, 200000),
    ('P003', 3, 'Premium internet with 200Mbps download speed', '200Mbps', 4320000, 800000),
    ('P004', 4, 'Ultra-fast 5G internet', '5G Unlimited', 5184000, 1200000),
    ('P005', 5, 'Home internet with 50Mbps speed', '50Mbps', 31536000, 350000);

INSERT INTO complaints (user_id, packet_internets_code, village, complaint_message, technician_note, reply, status, created_at)
VALUES
    (1, 'P001', 'Kaliwuluh', 'Internet connection is intermittent and slow', NULL, NULL, 'rejected', 1672531210000),
    (2, 'P002', 'Sendangadi', 'The internet speed is not matching the package description', NULL, NULL, 'accepted', 1672531211000),
    (3, 'P003', 'Pemecutan Kelod', 'My internet disconnects frequently', 'Technician visit scheduled', 'Issue resolved after visit', 'accepted', 1672531212000),
    (4, 'P004', 'Cipaganti', 'Unable to connect to the internet despite multiple attempts', 'No further issues found', 'User requested refund', 'rejected', 1672531213000),
    (5, 'P005', 'Ketintang', 'The internet speed is very slow compared to the advertised speed', 'Technician advised upgrading the plan', 'Issue resolved after plan upgrade', 'accepted', 1672531214000);

INSERT INTO subs_packets (user_id, packet_internets_code, lifetime, payment_time, status, created_at)
VALUES
    (1, 'P001', 1672531210000 + 2592000, 1672531215000, 'paid', 1672531210000),
    (2, 'P002', 1672531211000 + 8640000, 1672531216000, 'paid', 1672531211000),
    (3, 'P003', 1672531212000 + 4320000, 1672531217000, 'paid', 1672531212000),
    (4, 'P004', 1672531213000 + 5184000, 1672531218000, 'unpaid', 1672531213000),
    (5, 'P005', 1672531214000 + 31536000, 1672531219000, 'paid', 1672531214000);


SELECT * FROM users;
SELECT * FROM source_internets;
SELECT * FROM packet_internets;
SELECT * FROM complaints;
SELECT * FROM subs_packets;
