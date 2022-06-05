package src

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

func NaivePrepared(numberOfRows uint64, sqliteDB *sql.DB) {
	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Panicf(err.Error())
	}

	UserInsertionQuery := "insert into user(id,area,age,active) values(?,?,?,?)"

	stmt, err := tx.Prepare(UserInsertionQuery)

	if err != nil {
		fmt.Printf("Error in preparing statements: %s", err.Error())
	}

	for i := uint64(0); i < numberOfRows; i++ {
		newUser := &models.User{
			Id:     uint32(i + 1),
			Area:   GenerateRandomAreaCode(),
			Age:    GenerateRandomAge(),
			Active: GenerateRandomBooleanInt(),
		}
		stmt.Exec(newUser.Id, newUser.Area, newUser.Age, newUser.Active)
	}
	tx.Commit()

}
