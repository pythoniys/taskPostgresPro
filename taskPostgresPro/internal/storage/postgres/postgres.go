package postgres

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const fn = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS commands(
	    command_id int NOT NULL GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1),
	    command_name varchar(32) NOT NULL,
	    code_of_command text NOT NULL,
	    result_code varchar(32) NOT NULL);
	CREATE INDEX IF NOT EXISTS commands_name_idx ON commands(command_name);`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return &Storage{db: db}, nil
}
