package src

import (
	"fmt"
	"log"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/db"
	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

func Naive(numberOfRows uint64) {
	sqliteDB, err := db.NewDB("file:data/users.db")

	if err != nil {
		fmt.Errorf("Error while creating db instance %s", err)
	}

	err = db.RunMigrationScripts(sqliteDB)

	if err != nil {
		fmt.Printf("failed running migrations %s \n", err.Error())
	}

	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Fatalf(err.Error())
	}

	newUser := &models.User{
		Id:     1,
		Area:   "trichy",
		Age:    21,
		Active: 1,
	}

	UserInsertionQuery := "insert into user(id,area,age,active) values(?,?,?,?)"

	tx.Exec(UserInsertionQuery, newUser.Id, newUser.Area, newUser.Age, newUser.Active)

	UserSearchQuery := "select * from user"

	result, err := tx.Query(UserSearchQuery)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for result.Next() {
		var id, age, active uint32
		var area string
		if err := result.Scan(&id, &area, &age, &active); err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(id, area, age, active)
	}

}
