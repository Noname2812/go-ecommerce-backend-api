-- Transportation Init Data
-- This file contains sample data for routes, buses, seats, and trips
-- Run this after creating the tables

-- =====================================================
-- ROUTES DATA
-- =====================================================
INSERT INTO routes (route_start_location, route_end_location, route_estimated_duration) VALUES
('Hà Nội', 'Hồ Chí Minh', 1440), -- 24 hours
('Hà Nội', 'Đà Nẵng', 720),      -- 12 hours
('Hà Nội', 'Nha Trang', 1080),   -- 18 hours
('Hà Nội', 'Huế', 600),          -- 10 hours
('Hồ Chí Minh', 'Đà Nẵng', 720), -- 12 hours
('Hồ Chí Minh', 'Nha Trang', 360), -- 6 hours
('Hồ Chí Minh', 'Huế', 600),     -- 10 hours
('Đà Nẵng', 'Huế', 120),         -- 2 hours
('Đà Nẵng', 'Nha Trang', 360),   -- 6 hours
('Nha Trang', 'Huế', 480);       -- 8 hours

-- =====================================================
-- BUSES DATA
-- =====================================================
INSERT INTO buses (bus_license_plate, bus_company, bus_price, bus_capacity) VALUES
-- Luxury buses (45 seats)
('29A-12345', 'Phương Trang', 1500000.00, 45),
('29A-12346', 'Phương Trang', 1500000.00, 45),
('29A-12347', 'Phương Trang', 1500000.00, 45),
('29A-12348', 'Phương Trang', 1500000.00, 45),
('29A-12349', 'Phương Trang', 1500000.00, 45),

-- Standard buses (35 seats)
('29B-54321', 'Mai Linh', 1200000.00, 35),
('29B-54322', 'Mai Linh', 1200000.00, 35),
('29B-54323', 'Mai Linh', 1200000.00, 35),
('29B-54324', 'Mai Linh', 1200000.00, 35),
('29B-54325', 'Mai Linh', 1200000.00, 35),

-- Economy buses (30 seats)
('29C-98765', 'Thành Bưởi', 800000.00, 30),
('29C-98766', 'Thành Bưởi', 800000.00, 30),
('29C-98767', 'Thành Bưởi', 800000.00, 30),
('29C-98768', 'Thành Bưởi', 800000.00, 30),
('29C-98769', 'Thành Bưởi', 800000.00, 30),

-- VIP buses (20 seats)
('29D-11111', 'VIP Express', 2500000.00, 20),
('29D-11112', 'VIP Express', 2500000.00, 20),
('29D-11113', 'VIP Express', 2500000.00, 20),
('29D-11114', 'VIP Express', 2500000.00, 20),
('29D-11115', 'VIP Express', 2500000.00, 20);

-- =====================================================
-- SEATS DATA (for each bus)
-- =====================================================

-- Seats for Luxury buses (45 seats - 9 rows x 5 columns)
INSERT INTO seats (bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type) VALUES
-- Bus 1 (Luxury - 45 seats)
(1, 'A1', 1, 1, 1, 1), (1, 'A2', 1, 2, 1, 1), (1, 'A3', 1, 3, 1, 1), (1, 'A4', 1, 4, 1, 1), (1, 'A5', 1, 5, 1, 1),
(1, 'B1', 2, 1, 1, 1), (1, 'B2', 2, 2, 1, 1), (1, 'B3', 2, 3, 1, 1), (1, 'B4', 2, 4, 1, 1), (1, 'B5', 2, 5, 1, 1),
(1, 'C1', 3, 1, 1, 1), (1, 'C2', 3, 2, 1, 1), (1, 'C3', 3, 3, 1, 1), (1, 'C4', 3, 4, 1, 1), (1, 'C5', 3, 5, 1, 1),
(1, 'D1', 4, 1, 1, 1), (1, 'D2', 4, 2, 1, 1), (1, 'D3', 4, 3, 1, 1), (1, 'D4', 4, 4, 1, 1), (1, 'D5', 4, 5, 1, 1),
(1, 'E1', 5, 1, 1, 1), (1, 'E2', 5, 2, 1, 1), (1, 'E3', 5, 3, 1, 1), (1, 'E4', 5, 4, 1, 1), (1, 'E5', 5, 5, 1, 1),
(1, 'F1', 6, 1, 1, 1), (1, 'F2', 6, 2, 1, 1), (1, 'F3', 6, 3, 1, 1), (1, 'F4', 6, 4, 1, 1), (1, 'F5', 6, 5, 1, 1),
(1, 'G1', 7, 1, 1, 1), (1, 'G2', 7, 2, 1, 1), (1, 'G3', 7, 3, 1, 1), (1, 'G4', 7, 4, 1, 1), (1, 'G5', 7, 5, 1, 1),
(1, 'H1', 8, 1, 1, 1), (1, 'H2', 8, 2, 1, 1), (1, 'H3', 8, 3, 1, 1), (1, 'H4', 8, 4, 1, 1), (1, 'H5', 8, 5, 1, 1),
(1, 'I1', 9, 1, 1, 1), (1, 'I2', 9, 2, 1, 1), (1, 'I3', 9, 3, 1, 1), (1, 'I4', 9, 4, 1, 1), (1, 'I5', 9, 5, 1, 1),

-- Bus 2 (Luxury - 45 seats)
(2, 'A1', 1, 1, 1, 1), (2, 'A2', 1, 2, 1, 1), (2, 'A3', 1, 3, 1, 1), (2, 'A4', 1, 4, 1, 1), (2, 'A5', 1, 5, 1, 1),
(2, 'B1', 2, 1, 1, 1), (2, 'B2', 2, 2, 1, 1), (2, 'B3', 2, 3, 1, 1), (2, 'B4', 2, 4, 1, 1), (2, 'B5', 2, 5, 1, 1),
(2, 'C1', 3, 1, 1, 1), (2, 'C2', 3, 2, 1, 1), (2, 'C3', 3, 3, 1, 1), (2, 'C4', 3, 4, 1, 1), (2, 'C5', 3, 5, 1, 1),
(2, 'D1', 4, 1, 1, 1), (2, 'D2', 4, 2, 1, 1), (2, 'D3', 4, 3, 1, 1), (2, 'D4', 4, 4, 1, 1), (2, 'D5', 4, 5, 1, 1),
(2, 'E1', 5, 1, 1, 1), (2, 'E2', 5, 2, 1, 1), (2, 'E3', 5, 3, 1, 1), (2, 'E4', 5, 4, 1, 1), (2, 'E5', 5, 5, 1, 1),
(2, 'F1', 6, 1, 1, 1), (2, 'F2', 6, 2, 1, 1), (2, 'F3', 6, 3, 1, 1), (2, 'F4', 6, 4, 1, 1), (2, 'F5', 6, 5, 1, 1),
(2, 'G1', 7, 1, 1, 1), (2, 'G2', 7, 2, 1, 1), (2, 'G3', 7, 3, 1, 1), (2, 'G4', 7, 4, 1, 1), (2, 'G5', 7, 5, 1, 1),
(2, 'H1', 8, 1, 1, 1), (2, 'H2', 8, 2, 1, 1), (2, 'H3', 8, 3, 1, 1), (2, 'H4', 8, 4, 1, 1), (2, 'H5', 8, 5, 1, 1),
(2, 'I1', 9, 1, 1, 1), (2, 'I2', 9, 2, 1, 1), (2, 'I3', 9, 3, 1, 1), (2, 'I4', 9, 4, 1, 1), (2, 'I5', 9, 5, 1, 1),

-- Bus 3 (Luxury - 45 seats)
(3, 'A1', 1, 1, 1, 1), (3, 'A2', 1, 2, 1, 1), (3, 'A3', 1, 3, 1, 1), (3, 'A4', 1, 4, 1, 1), (3, 'A5', 1, 5, 1, 1),
(3, 'B1', 2, 1, 1, 1), (3, 'B2', 2, 2, 1, 1), (3, 'B3', 2, 3, 1, 1), (3, 'B4', 2, 4, 1, 1), (3, 'B5', 2, 5, 1, 1),
(3, 'C1', 3, 1, 1, 1), (3, 'C2', 3, 2, 1, 1), (3, 'C3', 3, 3, 1, 1), (3, 'C4', 3, 4, 1, 1), (3, 'C5', 3, 5, 1, 1),
(3, 'D1', 4, 1, 1, 1), (3, 'D2', 4, 2, 1, 1), (3, 'D3', 4, 3, 1, 1), (3, 'D4', 4, 4, 1, 1), (3, 'D5', 4, 5, 1, 1),
(3, 'E1', 5, 1, 1, 1), (3, 'E2', 5, 2, 1, 1), (3, 'E3', 5, 3, 1, 1), (3, 'E4', 5, 4, 1, 1), (3, 'E5', 5, 5, 1, 1),
(3, 'F1', 6, 1, 1, 1), (3, 'F2', 6, 2, 1, 1), (3, 'F3', 6, 3, 1, 1), (3, 'F4', 6, 4, 1, 1), (3, 'F5', 6, 5, 1, 1),
(3, 'G1', 7, 1, 1, 1), (3, 'G2', 7, 2, 1, 1), (3, 'G3', 7, 3, 1, 1), (3, 'G4', 7, 4, 1, 1), (3, 'G5', 7, 5, 1, 1),
(3, 'H1', 8, 1, 1, 1), (3, 'H2', 8, 2, 1, 1), (3, 'H3', 8, 3, 1, 1), (3, 'H4', 8, 4, 1, 1), (3, 'H5', 8, 5, 1, 1),
(3, 'I1', 9, 1, 1, 1), (3, 'I2', 9, 2, 1, 1), (3, 'I3', 9, 3, 1, 1), (3, 'I4', 9, 4, 1, 1), (3, 'I5', 9, 5, 1, 1),

-- Bus 4 (Luxury - 45 seats)
(4, 'A1', 1, 1, 1, 1), (4, 'A2', 1, 2, 1, 1), (4, 'A3', 1, 3, 1, 1), (4, 'A4', 1, 4, 1, 1), (4, 'A5', 1, 5, 1, 1),
(4, 'B1', 2, 1, 1, 1), (4, 'B2', 2, 2, 1, 1), (4, 'B3', 2, 3, 1, 1), (4, 'B4', 2, 4, 1, 1), (4, 'B5', 2, 5, 1, 1),
(4, 'C1', 3, 1, 1, 1), (4, 'C2', 3, 2, 1, 1), (4, 'C3', 3, 3, 1, 1), (4, 'C4', 3, 4, 1, 1), (4, 'C5', 3, 5, 1, 1),
(4, 'D1', 4, 1, 1, 1), (4, 'D2', 4, 2, 1, 1), (4, 'D3', 4, 3, 1, 1), (4, 'D4', 4, 4, 1, 1), (4, 'D5', 4, 5, 1, 1),
(4, 'E1', 5, 1, 1, 1), (4, 'E2', 5, 2, 1, 1), (4, 'E3', 5, 3, 1, 1), (4, 'E4', 5, 4, 1, 1), (4, 'E5', 5, 5, 1, 1),
(4, 'F1', 6, 1, 1, 1), (4, 'F2', 6, 2, 1, 1), (4, 'F3', 6, 3, 1, 1), (4, 'F4', 6, 4, 1, 1), (4, 'F5', 6, 5, 1, 1),
(4, 'G1', 7, 1, 1, 1), (4, 'G2', 7, 2, 1, 1), (4, 'G3', 7, 3, 1, 1), (4, 'G4', 7, 4, 1, 1), (4, 'G5', 7, 5, 1, 1),
(4, 'H1', 8, 1, 1, 1), (4, 'H2', 8, 2, 1, 1), (4, 'H3', 8, 3, 1, 1), (4, 'H4', 8, 4, 1, 1), (4, 'H5', 8, 5, 1, 1),
(4, 'I1', 9, 1, 1, 1), (4, 'I2', 9, 2, 1, 1), (4, 'I3', 9, 3, 1, 1), (4, 'I4', 9, 4, 1, 1), (4, 'I5', 9, 5, 1, 1),

-- Bus 5 (Luxury - 45 seats)
(5, 'A1', 1, 1, 1, 1), (5, 'A2', 1, 2, 1, 1), (5, 'A3', 1, 3, 1, 1), (5, 'A4', 1, 4, 1, 1), (5, 'A5', 1, 5, 1, 1),
(5, 'B1', 2, 1, 1, 1), (5, 'B2', 2, 2, 1, 1), (5, 'B3', 2, 3, 1, 1), (5, 'B4', 2, 4, 1, 1), (5, 'B5', 2, 5, 1, 1),
(5, 'C1', 3, 1, 1, 1), (5, 'C2', 3, 2, 1, 1), (5, 'C3', 3, 3, 1, 1), (5, 'C4', 3, 4, 1, 1), (5, 'C5', 3, 5, 1, 1),
(5, 'D1', 4, 1, 1, 1), (5, 'D2', 4, 2, 1, 1), (5, 'D3', 4, 3, 1, 1), (5, 'D4', 4, 4, 1, 1), (5, 'D5', 4, 5, 1, 1),
(5, 'E1', 5, 1, 1, 1), (5, 'E2', 5, 2, 1, 1), (5, 'E3', 5, 3, 1, 1), (5, 'E4', 5, 4, 1, 1), (5, 'E5', 5, 5, 1, 1),
(5, 'F1', 6, 1, 1, 1), (5, 'F2', 6, 2, 1, 1), (5, 'F3', 6, 3, 1, 1), (5, 'F4', 6, 4, 1, 1), (5, 'F5', 6, 5, 1, 1),
(5, 'G1', 7, 1, 1, 1), (5, 'G2', 7, 2, 1, 1), (5, 'G3', 7, 3, 1, 1), (5, 'G4', 7, 4, 1, 1), (5, 'G5', 7, 5, 1, 1),
(5, 'H1', 8, 1, 1, 1), (5, 'H2', 8, 2, 1, 1), (5, 'H3', 8, 3, 1, 1), (5, 'H4', 8, 4, 1, 1), (5, 'H5', 8, 5, 1, 1),
(5, 'I1', 9, 1, 1, 1), (5, 'I2', 9, 2, 1, 1), (5, 'I3', 9, 3, 1, 1), (5, 'I4', 9, 4, 1, 1), (5, 'I5', 9, 5, 1, 1);

-- Continue with Standard buses (35 seats - 7 rows x 5 columns)
INSERT INTO seats (bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type) VALUES
-- Bus 6 (Standard - 35 seats)
(6, 'A1', 1, 1, 1, 1), (6, 'A2', 1, 2, 1, 1), (6, 'A3', 1, 3, 1, 1), (6, 'A4', 1, 4, 1, 1), (6, 'A5', 1, 5, 1, 1),
(6, 'B1', 2, 1, 1, 1), (6, 'B2', 2, 2, 1, 1), (6, 'B3', 2, 3, 1, 1), (6, 'B4', 2, 4, 1, 1), (6, 'B5', 2, 5, 1, 1),
(6, 'C1', 3, 1, 1, 1), (6, 'C2', 3, 2, 1, 1), (6, 'C3', 3, 3, 1, 1), (6, 'C4', 3, 4, 1, 1), (6, 'C5', 3, 5, 1, 1),
(6, 'D1', 4, 1, 1, 1), (6, 'D2', 4, 2, 1, 1), (6, 'D3', 4, 3, 1, 1), (6, 'D4', 4, 4, 1, 1), (6, 'D5', 4, 5, 1, 1),
(6, 'E1', 5, 1, 1, 1), (6, 'E2', 5, 2, 1, 1), (6, 'E3', 5, 3, 1, 1), (6, 'E4', 5, 4, 1, 1), (6, 'E5', 5, 5, 1, 1),
(6, 'F1', 6, 1, 1, 1), (6, 'F2', 6, 2, 1, 1), (6, 'F3', 6, 3, 1, 1), (6, 'F4', 6, 4, 1, 1), (6, 'F5', 6, 5, 1, 1),
(6, 'G1', 7, 1, 1, 1), (6, 'G2', 7, 2, 1, 1), (6, 'G3', 7, 3, 1, 1), (6, 'G4', 7, 4, 1, 1), (6, 'G5', 7, 5, 1, 1);

-- Continue with remaining buses (simplified for brevity)
-- Bus 7-10 (Standard - 35 seats each)
INSERT INTO seats (bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type)
SELECT 
    bus_id,
    CONCAT(CHAR(65 + (seat_id - 1) DIV 5), ((seat_id - 1) % 5) + 1) as seat_number,
    ((seat_id - 1) DIV 5) + 1 as seat_row_no,
    ((seat_id - 1) % 5) + 1 as seat_column_no,
    1 as seat_floor_no,
    1 as seat_type
FROM (
    SELECT 7 as bus_id UNION SELECT 8 UNION SELECT 9 UNION SELECT 10
) buses
CROSS JOIN (
    SELECT 1 as seat_id UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION
    SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10 UNION
    SELECT 11 UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15 UNION
    SELECT 16 UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20 UNION
    SELECT 21 UNION SELECT 22 UNION SELECT 23 UNION SELECT 24 UNION SELECT 25 UNION
    SELECT 26 UNION SELECT 27 UNION SELECT 28 UNION SELECT 29 UNION SELECT 30 UNION
    SELECT 31 UNION SELECT 32 UNION SELECT 33 UNION SELECT 34 UNION SELECT 35
) seats;

-- Economy buses (30 seats - 6 rows x 5 columns)
INSERT INTO seats (bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type)
SELECT 
    bus_id,
    CONCAT(CHAR(65 + (seat_id - 1) DIV 5), ((seat_id - 1) % 5) + 1) as seat_number,
    ((seat_id - 1) DIV 5) + 1 as seat_row_no,
    ((seat_id - 1) % 5) + 1 as seat_column_no,
    1 as seat_floor_no,
    1 as seat_type
FROM (
    SELECT 11 as bus_id UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15
) buses
CROSS JOIN (
    SELECT 1 as seat_id UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION
    SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10 UNION
    SELECT 11 UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15 UNION
    SELECT 16 UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20 UNION
    SELECT 21 UNION SELECT 22 UNION SELECT 23 UNION SELECT 24 UNION SELECT 25 UNION
    SELECT 26 UNION SELECT 27 UNION SELECT 28 UNION SELECT 29 UNION SELECT 30
) seats;

-- VIP buses (20 seats - 4 rows x 5 columns)
INSERT INTO seats (bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type)
SELECT 
    bus_id,
    CONCAT(CHAR(65 + (seat_id - 1) DIV 5), ((seat_id - 1) % 5) + 1) as seat_number,
    ((seat_id - 1) DIV 5) + 1 as seat_row_no,
    ((seat_id - 1) % 5) + 1 as seat_column_no,
    1 as seat_floor_no,
    2 as seat_type -- VIP seats
FROM (
    SELECT 16 as bus_id UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20
) buses
CROSS JOIN (
    SELECT 1 as seat_id UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION
    SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9 UNION SELECT 10 UNION
    SELECT 11 UNION SELECT 12 UNION SELECT 13 UNION SELECT 14 UNION SELECT 15 UNION
    SELECT 16 UNION SELECT 17 UNION SELECT 18 UNION SELECT 19 UNION SELECT 20
) seats;

-- =====================================================
-- TRIPS DATA
-- =====================================================
INSERT INTO trips (route_id, bus_id, trip_departure_time, trip_arrival_time, trip_base_price) VALUES
-- Hà Nội - Hồ Chí Minh routes
(1, 1, '2025-01-15 18:00:00', '2025-01-16 18:00:00', 1500000.00),
(1, 2, '2025-01-16 18:00:00', '2025-01-17 18:00:00', 1500000.00),
(1, 3, '2025-01-17 18:00:00', '2025-01-18 18:00:00', 1500000.00),

-- Hà Nội - Đà Nẵng routes
(2, 4, '2025-01-15 20:00:00', '2025-01-16 08:00:00', 800000.00),
(2, 5, '2025-01-16 20:00:00', '2025-01-17 08:00:00', 800000.00),
(2, 6, '2025-01-17 20:00:00', '2025-01-18 08:00:00', 800000.00),

-- Hà Nội - Nha Trang routes
(3, 7, '2025-01-15 19:00:00', '2025-01-16 13:00:00', 1200000.00),
(3, 8, '2025-01-16 19:00:00', '2025-01-17 13:00:00', 1200000.00),

-- Hà Nội - Huế routes
(4, 9, '2025-01-15 21:00:00', '2025-01-16 07:00:00', 600000.00),
(4, 10, '2025-01-16 21:00:00', '2025-01-17 07:00:00', 600000.00),

-- Hồ Chí Minh - Đà Nẵng routes
(5, 11, '2025-01-15 18:00:00', '2025-01-16 06:00:00', 800000.00),
(5, 12, '2025-01-16 18:00:00', '2025-01-17 06:00:00', 800000.00),

-- Hồ Chí Minh - Nha Trang routes
(6, 13, '2025-01-15 08:00:00', '2025-01-15 14:00:00', 400000.00),
(6, 14, '2025-01-16 08:00:00', '2025-01-16 14:00:00', 400000.00),

-- Hồ Chí Minh - Huế routes
(7, 15, '2025-01-15 20:00:00', '2025-01-16 06:00:00', 600000.00),
(7, 16, '2025-01-16 20:00:00', '2025-01-17 06:00:00', 600000.00),

-- Đà Nẵng - Huế routes
(8, 17, '2025-01-15 07:00:00', '2025-01-15 09:00:00', 150000.00),
(8, 18, '2025-01-16 07:00:00', '2025-01-16 09:00:00', 150000.00),

-- Đà Nẵng - Nha Trang routes
(9, 19, '2025-01-15 14:00:00', '2025-01-15 20:00:00', 400000.00),
(9, 20, '2025-01-16 14:00:00', '2025-01-16 20:00:00', 400000.00),

-- Nha Trang - Huế routes
(10, 1, '2025-01-15 06:00:00', '2025-01-15 14:00:00', 500000.00),
(10, 2, '2025-01-16 06:00:00', '2025-01-16 14:00:00', 500000.00);

-- =====================================================
-- SAMPLE TRIP SEAT LOCKS (for testing)
-- =====================================================
INSERT INTO trip_seat_locks (trip_id, seat_id, locked_by_booking_id, trip_seat_lock_status, trip_seat_lock_expires_at) VALUES
-- Trip 1: Some seats are locked/booked
(1, 1, 1001, 2, DATE_ADD(NOW(), INTERVAL 30 MINUTE)), -- Locked
(1, 2, 1002, 3, NULL), -- Booked
(1, 3, NULL, 1, NULL), -- Available
(1, 4, NULL, 1, NULL), -- Available
(1, 5, 1003, 2, DATE_ADD(NOW(), INTERVAL 15 MINUTE)), -- Locked

-- Trip 2: Most seats available
(2, 1, NULL, 1, NULL), -- Available
(2, 2, NULL, 1, NULL), -- Available
(2, 3, 1004, 3, NULL), -- Booked
(2, 4, NULL, 1, NULL), -- Available
(2, 5, NULL, 1, NULL); -- Available 