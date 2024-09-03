CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    gender VARCHAR(50),
    birthday DATE NOT NULL,
    launchpad_id VARCHAR(255) NOT NULL,
    destination_id INT NOT NULL,
    launch_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);