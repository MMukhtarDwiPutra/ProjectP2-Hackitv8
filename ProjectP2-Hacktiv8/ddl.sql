CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password VARCHAR(255) NOT NULL,
    balance NUMERIC(10, 2) NOT NULL
);

CREATE TABLE rooms (
    room_id SERIAL PRIMARY KEY,
    price VARCHAR(255) NOT NULL,
    room_type VARCHAR(128) NOT NULL
);

CREATE TABLE bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    room_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (room_id) REFERENCES rooms(room_id)
);

INSERT INTO users (full_name, email, password, balance) 
VALUES 
('John Doe', 'john.doe@example.com', 'password123', 1000.00),
('Jane Smith', 'jane.smith@example.com', 'password456', 1500.50);

INSERT INTO rooms (price, room_type) 
VALUES 
('200.00', 'Deluxe'),
('150.00', 'Standard'),
('300.00', 'Suite');

INSERT INTO bookings (user_id, room_id)
VALUES 
(1, 2), -- John Doe booked a Standard room
(2, 1), -- Jane Smith booked a Deluxe room
(1, 3); -- John Doe booked a Suite room