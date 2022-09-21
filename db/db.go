package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/qna-page/qna-page/resource/user"
)

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}

	// Create tables
	db.MustExec(user.UserSchema)
	// db.MustExec(variable.VariableSchema)

	fmt.Println("Successfully connected to DB!")

	return db
}
