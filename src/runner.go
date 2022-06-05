package src

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type InsertionMethod uint32

const (
	M_Naive InsertionMethod = iota
	M_Naive_Async
	// add more methods here
)

func (i InsertionMethod) ToString() string {
	switch i {
	case M_Naive:
		return "M_Naive"

	case M_Naive_Async:
		return "M_Naive_Async"
	default:
		return "NULL"
	}
}

func Runner(method InsertionMethod, numberOfRows uint64, sqliteDB *sql.DB) {
	switch method {
	case M_Naive:
		{
			start := time.Now()
			Naive(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	case M_Naive_Async:
		{
			start := time.Now()
			NaiveAsync(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	}
	// validate data
	if !ValidateTable(numberOfRows, sqliteDB) {
		log.Fatalf("Rows not inserted properly")
	}

	// truncate table
	err := TruncateUserTable(sqliteDB)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
