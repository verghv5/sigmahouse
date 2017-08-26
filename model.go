package sigmahouse

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
	Date string `json:"date"`
}

func (i *issue) getIssue(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (i *issue) updateIssue(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (i *issue) deleteIssue(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (i *issue) createIssue(db *sql.DB) error {
	return errors.New("Not implemented")
}


func getIssues(db *sql.DB, start, count int) ([]issue, error) {
	return nil, errors.New("Not implemented")
}

