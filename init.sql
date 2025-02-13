CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    room_number VARCHAR(255),
    is_authorised BOOLEAN DEFAULT FALSE
);
