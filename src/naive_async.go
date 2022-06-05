package src

import (
	"database/sql"
	"log"
	"sync"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

// naive + async using go-routines
func NaiveAsync(numberOfRows uint64, sqliteDB *sql.DB) {
	tx, err := sqliteDB.Begin()

	// wait groups
	var wg sync.WaitGroup

	if err != nil {
		log.Panicf(err.Error())
	}
	UserInsertionQuery := "insert into user(id,area,age,active) values(?,?,?,?)"

	for i := uint64(0); i < numberOfRows; i++ {
		wg.Add(1)
		go func(id uint64) {
			defer wg.Done()
			newUser := &models.User{
				Id:     uint32(id + 1),
				Area:   GenerateRandomAreaCode(),
				Age:    GenerateRandomAge(),
				Active: GenerateRandomBooleanInt(),
			}
			tx.Exec(UserInsertionQuery, newUser.Id, newUser.Area, newUser.Age, newUser.Active)
		}(i)
	}
	wg.Wait()
	tx.Commit()
}
