-- Users Table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,                          -- Auto-increment primary key
    full_name VARCHAR(255) NOT NULL,                    -- User's full name
    email VARCHAR(128) NOT NULL UNIQUE,                 -- Email must be unique
    password VARCHAR(255) NOT NULL,                     -- User's hashed password
    balance NUMERIC(10, 2) NOT NULL DEFAULT 0.00,       -- Balance with precision
    jwt_token VARCHAR(255),                             -- JWT token, nullable until generated
    is_activated VARCHAR(20) NOT NULL DEFAULT 'NOT YET' 
);

-- Rooms Table
CREATE TABLE rooms (
    room_id SERIAL PRIMARY KEY,                          -- Auto-increment primary key
    price NUMERIC(10, 2) NOT NULL,                       -- Price as numeric for calculations
    room_type VARCHAR(128) NOT NULL,                    -- Room type (e.g., Single, Double)
    availability_status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE' 
);

-- Bookings Table
CREATE TABLE bookings (
    booking_id SERIAL PRIMARY KEY,                      -- Auto-increment primary key
    user_id INT NOT NULL,                               -- Foreign key to users
    room_id INT NOT NULL,                               -- Foreign key to rooms
    booking_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Timestamp of booking
    date_in DATE NOT NULL,
    date_out DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);

CREATE TABLE rents (
    rent_id SERIAL PRIMARY KEY,         -- Auto-increment primary key
    user_id INT NOT NULL,               -- Foreign key to users
    room_id INT NOT NULL,               -- Foreign key to rooms
    date_in DATE NOT NULL,              -- Rent start date
    date_out DATE NOT NULL,             -- Rent end date
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms(room_id) ON DELETE CASCADE
);

-- INSERT INTO users (full_name, email, password, balance) 
-- VALUES 
-- ('Muhammad Mukhtar Dwi Putra', 'putra16.mp@gmail.com', '$2a$10$m8NCL7NYth/H65qPH3l0feu3pQ.ww8YGjxo3HbCXwLGokSgd/bB9G', 0.00);

INSERT INTO rooms (price, room_type) 
VALUES 
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1200.00', 'Standard'),
('1500.00', 'Deluxe'),
('2000.00', 'Suite'),
('2500.00', 'Excecutive Suite');

-- INSERT INTO bookings (user_id, room_id)
-- VALUES 
-- (1, 2), -- John Doe booked a Standard room
-- (2, 1), -- Jane Smith booked a Deluxe room
-- (1, 3); -- John Doe booked a Suite room