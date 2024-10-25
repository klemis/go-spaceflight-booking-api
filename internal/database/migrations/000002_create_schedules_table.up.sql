CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    launchpad_id VARCHAR(255) NOT NULL,
    destination_id INT NOT NULL,
    day_of_week INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_launchpad_day UNIQUE (launchpad_id, day_of_week)
);