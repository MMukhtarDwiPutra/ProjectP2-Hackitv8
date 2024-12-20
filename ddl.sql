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

CREATE TABLE webhook_xendit_payments(
    id SERIAL PRIMARY KEY,
    invoice_id VARCHAR(255) NOT NULL,
    user_id_app INT NOT NULL,
    status VARCHAR(20) NOT NULL
);

INSERT INTO public.webhook_xendit_payments VALUES (2, 'INV1', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (4, 'INV3', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (6, 'INV5', 4, 'PAID');
INSERT INTO public.webhook_xendit_payments VALUES (8, 'INV7', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (10, 'INV9', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (12, 'INV11', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (14, 'INV13', 4, 'PENDING');
INSERT INTO public.webhook_xendit_payments VALUES (16, 'INV15', 4, 'PENDING');

INSERT INTO public.rooms VALUES (2, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (3, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (4, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (5, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (6, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (7, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (8, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (9, 1200.00, 'Standard', 'Available');
INSERT INTO public.rooms VALUES (10, 1500.00, 'Deluxe', 'Available');
INSERT INTO public.rooms VALUES (11, 2000.00, 'Suite', 'Available');
INSERT INTO public.rooms VALUES (12, 2500.00, 'Excecutive Suite', 'Available');
INSERT INTO public.rooms VALUES (1, 1200.00, 'Standard', 'Available');

INSERT INTO public.users VALUES (6, 'Muhammad Mukhtar Dwi Putra', 'screedmp@gmail.com', '$2a$10$XofwmXQWi79QG5DTWuTXgOtHbJn8mXnYZ0ueJIjNf7WqlgmQF80D.', 0.00, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNjcmVlZG1wQGdtYWlsLmNvbSIsImV4cCI6MTczNDg4ODIzMn0.O4ioCo_rKcjAUJGNiS5ZPr8fAQDPl115HkXiTyuH1Rg', 'Activated');
INSERT INTO public.users VALUES (4, 'Muhammad Mukhtar Dwi Putra', 'putra16.mp@gmail.com', '$2a$10$0gtyZC87Qv.uCgEjX5Q9B.P.H7t.l8ZOp2g2t1G57ze6COsopSH26', 28600.00, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InB1dHJhMTYubXBAZ21haWwuY29tIiwiZXhwIjoxNzM0ODQxNDI1fQ.N0KV7iw1kWrI1Bg_JapR5yH9sAtvc8epobStvh51Lw4', 'Activated');

INSERT INTO public.bookings VALUES (1, 4, 1, '2024-12-20 11:19:24.472816', '2024-12-20', '2024-12-30');
INSERT INTO public.bookings VALUES (2, 4, 1, '2024-12-20 14:45:24.388399', '2024-12-20', '2024-12-30');