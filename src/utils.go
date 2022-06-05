package src

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)

// generates a random area code of length 6
func GenerateRandomAreaCode() string {
	return fmt.Sprint((rand.Intn(999999)))
}

func GenerateRandomAge() uint32 {
	return uint32((rand.Int31n(80)))
}

func GenerateRandomBooleanInt() uint32 {
	return uint32((rand.Int31n(2)))
}

func TruncateUserTable(sqliteDB *sql.DB) error {
	truncateQuery := "truncate user"
	tx, err := sqliteDB.Begin()

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	tx.Exec(truncateQuery)

	tx.Commit()
	return nil
}

// checks if number of rows given are present in the table
func ValidateTable(numberOfRows uint64, sqliteDB *sql.DB) bool {
	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Fatalf(err.Error())
	}

	CountRowsQuery := "select count(*) from user"
	res, err := tx.Query(CountRowsQuery)

	if err != nil {
		log.Fatalf(err.Error())
	}

	var nrows uint64

	res.Next()

	if err := res.Scan(&nrows); err != nil {
		log.Fatalf(err.Error())
	}

	return numberOfRows == nrows
}
