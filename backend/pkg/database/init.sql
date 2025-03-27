CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    home_team VARCHAR(255) NOT NULL,
    away_team VARCHAR(255) NOT NULL,
    match_date DATE NOT NULL,
    goals INTEGER DEFAULT 0,
    yellow_cards INTEGER DEFAULT 0,
    red_cards INTEGER DEFAULT 0,
    extra_time BOOLEAN DEFAULT FALSE
);


INSERT INTO matches (home_team, away_team, match_date, goals, yellow_cards, red_cards, extra_time)
VALUES 
('Real Madrid', 'Barcelona', '2025-04-01', 2, 3, 0, FALSE),
('Atlético Madrid', 'Sevilla', '2025-04-02', 1, 1, 1, FALSE),
('Valencia', 'Getafe', '2025-04-03', 3, 0, 0, TRUE);
