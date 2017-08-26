package main

import (
	"database/sql"
	"errors"
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
	return db.QueryRow("SELECT title, description FROM sigmahouse WHERE id=$1",
		i.ID).Scan(&i.Title, &i.Description, &i.Priority, &i.ReportDate)
}

func (i *issue) updateIssue(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (i *issue) deleteIssue(db *sql.DB) error {
	return errors.New("Not implemented")
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

//func getIssues(db *sql.DB, start, count int) ([]issue, error) {
//	return nil, errors.New("Not implemented")
//}

