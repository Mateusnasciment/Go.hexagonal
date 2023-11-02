package entity

import (
	"database/sql"
	"time"
)

type Entity struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Entity) Create() error {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO entities(name, created_at, updated_at) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return err
	}

	e.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
