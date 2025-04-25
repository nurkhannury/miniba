// db/database.go

package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type ReportRecord struct {
	ID         int    `json:"id"`
	Report     string `json:"report"`
	ReportDate string `json:"report_date"`
}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./miniba.db")
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	_, err = DB.Exec(`
      CREATE TABLE IF NOT EXISTS reports (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         email TEXT NOT NULL,
         password TEXT NOT NULL,
         report TEXT NOT NULL,
         report_date DATETIME DEFAULT CURRENT_TIMESTAMP
      );
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveReport(email, password, report string) error {
	_, err := DB.Exec(
		`INSERT INTO reports(email,password,report,report_date)
           VALUES(?,?,?,?)`,
		email, password, report,
		time.Now().Format("2006-01-02 15:04:05"),
	)
	return err
}

func GetReportsByEmailAndPassword(email, password string) ([]ReportRecord, error) {
	rows, err := DB.Query(
		`SELECT id,report,report_date
           FROM reports
          WHERE email=? AND password=?`,
		email, password,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recs []ReportRecord
	for rows.Next() {
		var r ReportRecord
		if err := rows.Scan(&r.ID, &r.Report, &r.ReportDate); err != nil {
			return nil, err
		}
		recs = append(recs, r)
	}
	return recs, nil
}
