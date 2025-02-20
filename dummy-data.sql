INSERT INTO users (NIK, full_name, province, district, sub_district, village, coordinate, roles, username, password, created_at, updated_at, deleted_at)
VALUES
    ('1234567890123456', 'John Doe', 'Jawa Barat', 'Bandung', 'Cidadap', 'Cimahi', ST_GeomFromText('POINT(107.6191 -6.9147)'), 'admin', 'johndoe', 'password123', 1632153600000, 1632157200000, NULL),
    ('2345678901234567', 'Jane Smith', 'Banten', 'Serang', 'Kasemen', 'Kramatwatu', ST_GeomFromText('POINT(106.1504 -6.1150)'), 'technician', 'janesmith', 'password456', 1632240000000, 1632243600000, NULL),
    ('3456789012345678', 'Alice Johnson', 'DKI Jakarta', 'Jakarta Selatan', 'Pondok Indah', 'Pondok Indah', ST_GeomFromText('POINT(106.7944 -6.2617)'), 'region', 'alicejohnson', 'password789', 1632326400000, 1632330000000, NULL),
    ('4567890123456789', 'Bob Brown', 'Jawa Tengah', 'Semarang', 'Pedurungan', 'Tlogosari', ST_GeomFromText('POINT(110.4387 -7.0027)'), 'citizens-data', 'bobbrown', 'password101', 1632412800000, 1632416400000, NULL),
    ('5678901234567890', 'Charlie Davis', 'Yogyakarta', 'Sleman', 'Mlati', 'Depok', ST_GeomFromText('POINT(110.4143 -7.7442)'), 'admin', 'charliedavis', 'password202', 1632499200000, 1632502800000, NULL),
    ('6789012345678901', 'David Miller', 'Jawa Timur', 'Surabaya', 'Gubeng', 'Bubutan', ST_GeomFromText('POINT(112.7492 -7.2491)'), 'technician', 'davidmiller', 'password303', 1632585600000, 1632589200000, NULL);

INSERT INTO packet_internets (code, provider_name, source, packet, duration, price)
VALUES
    ('PKT001', 'Provider A', 'Fiber Optic', '100Mbps', 31536000000, 100000),
    ('PKT002', 'Provider B', 'Satellite', '50Mbps', 25920000000, 80000),
    ('PKT003', 'Provider C', 'Fiber Optic', '200Mbps', 31536000000, 150000),
    ('PKT004', 'Provider D', '5G', '150Mbps', 18144000000, 120000),
    ('PKT005', 'Provider E', 'Wi-Fi', '25Mbps', 6307200000, 50000),
    ('PKT006', 'Provider F', 'Fiber Optic', '500Mbps', 31536000000, 200000);


INSERT INTO complaints (user_id, packet_internets_code, village, complaint_message, technician_note, reply, status, created_at, updated_at, deleted_at)
VALUES
    (1, 'PKT001', 'Cimahi', 'Connection speed is too slow', 'Checked the cable, no issues', 'Speed was restored', 'accepted', 1632153600000, 1632157200000, NULL),
    (2, 'PKT002', 'Kramatwatu', 'No internet signal', 'Technician dispatched', 'Signal restored', 'accepted', 1632240000000, 1632243600000, NULL),
    (3, 'PKT003', 'Pondok Indah', 'Frequent disconnections', 'Router firmware update done', 'Problem solved', 'accepted', 1632326400000, 1632330000000, NULL),
    (4, 'PKT004', 'Tlogosari', 'Packet expired before end of term', 'Investigation completed', 'Refund issued', 'accepted', 1632412800000, 1632416400000, NULL),
    (5, 'PKT005', 'Depok', 'Slow upload speed', 'Checked the modem, working fine', 'Escalated to senior technician', 'rejected', 1632499200000, 1632502800000, NULL),
    (6, 'PKT006', 'Bubutan', 'No connection for 2 days', 'Replaced the router', 'Service restored', 'accepted', 1632585600000, 1632589200000, NULL);

INSERT INTO subs_packet (user_id, packet_internets_code, lifetime, payment_time, status, created_at, updated_at, deleted_at)
VALUES
    (1, 'PKT001', 1635724800000, 1632153600000, 'paid', 1632153600000, 1632157200000, NULL),
    (2, 'PKT002', 1625097600000, 1632240000000, 'paid', 1632240000000, 1632243600000, NULL),
    (3, 'PKT003', 1640995200000, 1632326400000, 'unpaid', 1632326400000, 1632330000000, NULL),
    (4, 'PKT004', 1612137600000, 1632412800000, 'paid', 1632412800000, 1632416400000, NULL),
    (5, 'PKT005', 1609459200000, 1632499200000, 'unpaid', 1632499200000, 1632502800000, NULL),
    (6, 'PKT006', 1643673600000, 1632585600000, 'paid', 1632585600000, 1632589200000, NULL);
