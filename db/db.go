package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	id	VARCHAR(250) PRIMARY KEY,
    name	VARCHAR(250)  DEFAULT '',
	email	VARCHAR(250) DEFAULT '',
	password	VARCHAR(250) DEFAULT ''
);
INSERT INTO user (id,name,email,password) VALUES('EbXKRm6MuCOqa18j0Mqcx','example','example@example.com', '');
`

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}

	db.MustExec(schema)

	fmt.Println("Successfully connected to DB!")

	return db
}
