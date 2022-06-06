package src

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

func BatchedCheck(sqliteDB *sql.DB) {
	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Panicf(err.Error())
	}

	queryString := make([]string, 0, 3)
	queryArgs := make([]interface{}, 0, 3*4)

	for i := uint64(0); i < 3; i++ {
		queryString = append(queryString, "(?,?,?,?)")
		newUser := &models.User{
			Id:     uint32(i + 1),
			Area:   GenerateRandomAreaCode(),
			Age:    GenerateRandomAge(),
			Active: GenerateRandomBooleanInt(),
		}
		queryArgs = append(queryArgs, newUser.Id)
		queryArgs = append(queryArgs, newUser.Area)
		queryArgs = append(queryArgs, newUser.Age)
		queryArgs = append(queryArgs, newUser.Active)
	}

	query := fmt.Sprintf("insert into user(id,area,age,active) values %s", strings.Join(queryString, ","))

	stmt, err := tx.Prepare(query)

	if err != nil {
		log.Panicf(err.Error())
	}

	stmt.Exec(queryArgs...)

	res, err := tx.Query("select * from user")

	if err != nil {
		log.Panicf(err.Error())
	}

	for res.Next() {
		var id, age, active uint64
		var area string
		res.Scan(&id, &area, &age, &active)
		fmt.Println(id, area, age, active)
	}

}
