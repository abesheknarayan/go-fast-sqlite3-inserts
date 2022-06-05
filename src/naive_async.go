package src

import (
	"database/sql"
	"log"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

func NaiveAsync(numberOfRows uint64, sqliteDB *sql.DB) {
	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Fatalf(err.Error())
	}
	UserInsertionQuery := "insert into user(id,area,age,active) values(?,?,?,?)"

	for i := uint64(0); i < numberOfRows; i++ {

		go func(id uint64) {
			newUser := &models.User{
				Id:     uint32(id + 1),
				Area:   GenerateRandomAreaCode(),
				Age:    GenerateRandomAge(),
				Active: GenerateRandomBooleanInt(),
			}
			tx.Exec(UserInsertionQuery, newUser.Id, newUser.Area, newUser.Age, newUser.Active)
			tx.Commit()
		}(i)
	}
}
