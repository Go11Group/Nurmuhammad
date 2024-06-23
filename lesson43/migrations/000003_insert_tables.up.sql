-- Insert data into users table
INSERT INTO users (name, phone, age) VALUES 
('Alice Johnson', '555-1234', 30),
('Bob Smith', '555-5678', 45),
('Carol White', '555-8765', 22);

-- Insert data into card table
INSERT INTO card (number, user_id) VALUES 
('1234567890123456', 1),
('2345678901234567', 2),
('3456789012345678', 3);

-- Insert data into station table
INSERT INTO station (name) VALUES 
('Station 1'),
('Station 2'),
('Station 3');

-- Insert data into terminal table
INSERT INTO terminal (id, station_id) VALUES 
('550e8400-e29b-41d4-a716-446655440000', 1),
('550e8400-e29b-41d4-a716-446655440001', 2),
('550e8400-e29b-41d4-a716-446655440002', 3);

-- Insert data into transaction table
INSERT INTO transaction (card_id, amount, terminal_id, transaction_type) VALUES 
(1, 100.00, '550e8400-e29b-41d4-a716-446655440000', 'credit'),
(2, 50.00, '550e8400-e29b-41d4-a716-446655440001', 'debit'),
(3, 75.50, '550e8400-e29b-41d4-a716-446655440002', 'credit');
