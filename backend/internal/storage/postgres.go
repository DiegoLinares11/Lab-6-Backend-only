package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DiegoLinares11/Lab-6-Backend-Only/internal/models"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{db: db}
}

func (ps *PostgresStorage) GetAllMatches() ([]models.Match, error) {
	query := "SELECT id, home_team, away_team, match_date, goals, yellow_cards, red_cards, extra_time FROM matches"
	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var m models.Match
		err := rows.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate, &m.Goals, &m.YellowCards, &m.RedCards, &m.ExtraTime)
		if err != nil {
			return nil, err
		}
		matches = append(matches, m)
	}

	return matches, nil
}

func (ps *PostgresStorage) GetMatchByID(id int) (*models.Match, error) {
	query := "SELECT id, home_team, away_team, match_date, goals, yellow_cards, red_cards, extra_time FROM matches WHERE id = $1"
	row := ps.db.QueryRow(query, id)

	var m models.Match
	err := row.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate, &m.Goals, &m.YellowCards, &m.RedCards, &m.ExtraTime)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("match not found")
	}
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (ps *PostgresStorage) CreateMatch(match *models.Match) error {
	// Convertir la fecha de string a time.Time
	matchDate, err := time.Parse("2006-01-02", match.MatchDate)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	query := `
        INSERT INTO matches (home_team, away_team, match_date) 
        VALUES ($1, $2, $3) RETURNING id
    `
	err = ps.db.QueryRow(query, match.HomeTeam, match.AwayTeam, matchDate).Scan(&match.ID)
	return err
}

func (ps *PostgresStorage) UpdateMatch(match *models.Match) error {
	// Convertir la fecha de string a time.Time
	matchDate, err := time.Parse("2006-01-02", match.MatchDate)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	query := `
        UPDATE matches 
        SET home_team = $2, away_team = $3, match_date = $4
        WHERE id = $1
    `
	_, err = ps.db.Exec(query, match.ID, match.HomeTeam, match.AwayTeam, matchDate)
	return err
}

func (ps *PostgresStorage) DeleteMatch(id int) error {
	query := "DELETE FROM matches WHERE id = $1"
	_, err := ps.db.Exec(query, id)
	return err
}

func (ps *PostgresStorage) RegisterGoal(id int) error {
	query := "UPDATE matches SET goals = goals + 1 WHERE id = $1"
	_, err := ps.db.Exec(query, id)
	return err
}

func (ps *PostgresStorage) RegisterYellowCard(id int) error {
	query := "UPDATE matches SET yellow_cards = yellow_cards + 1 WHERE id = $1"
	_, err := ps.db.Exec(query, id)
	return err
}

func (ps *PostgresStorage) RegisterRedCard(id int) error {
	query := "UPDATE matches SET red_cards = red_cards + 1 WHERE id = $1"
	_, err := ps.db.Exec(query, id)
	return err
}

func (ps *PostgresStorage) SetExtraTime(id int) error {
	query := "UPDATE matches SET extra_time = true WHERE id = $1"
	_, err := ps.db.Exec(query, id)
	return err
}
