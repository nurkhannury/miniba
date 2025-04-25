package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./miniba.db")
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	// создаем таблицу, если не существует
	createTable := `
	CREATE TABLE IF NOT EXISTS reports (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		report TEXT NOT NULL,
		report_date DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

// SaveReport сохраняет отчёт в базу с паролем
func SaveReport(email, password, report string) error {
	_, err := DB.Exec(
		`INSERT INTO reports (email, password, report, report_date)
		 VALUES (?, ?, ?, ?)`,
		email, password, report, time.Now().Format("2006-01-02 15:04:05"),
	)
	return err
}

// GetReportsByEmailAndPassword returns all reports for a given email and password
func GetReportsByEmailAndPassword(email, password string) ([]string, error) {
	rows, err := DB.Query("SELECT report FROM reports WHERE email = ? AND password = ?", email, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []string
	for rows.Next() {
		var report string
		if err := rows.Scan(&report); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
