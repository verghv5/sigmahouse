package main

import (
	"database/sql"
)

/*
 model.go contains the structures used for api requests
 */

type issue struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Priority string `json:"priority"`
	ReportDate string `json:"reportdate"`
}

func (i *issue) getIssue(db *sql.DB) error {
	return db.QueryRow("SELECT title, description, priority, reportdate FROM issues WHERE id=$1",
		i.ID).Scan(&i.Title, &i.Description, &i.Priority, &i.ReportDate)
}

func (i *issue) updateIssue(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE issues SET title=$1, description=$2, priority=$3, reportdate=$4 WHERE id=$5",
			i.Title, i.Description, i.Priority, i.ReportDate, i.ID)

	return err
}

func (i *issue) deleteIssue(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM issues WHERE id=$1", i.ID)

	return err
}

func (i *issue) createIssue(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO issues(title, description, priority, reportdate) VALUES($1, $2, $3, $4) RETURNING id",
		i.Title, i.Description, i.Priority, i.ReportDate).Scan(&i.ID)

	if err != nil {
		return err
	}

	return nil
}

func getIssues(db *sql.DB) (map[string][]issue, error) {
	rows, err := db.Query("SELECT id, title, description, priority, reportdate FROM issues ORDER BY priority ASC, reportdate DESC LIMIT 50")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	issues := []issue{}

	for rows.Next() {
		var i issue
		if err := rows.Scan(&i.ID, &i.Title, &i.Description, &i.Priority, &i.ReportDate); err != nil {
			return nil, err
		}
		issues = append(issues, i)
	}

	jsonObj := map[string][]issue{"sigmaIssues": issues}

	return jsonObj, nil
}

