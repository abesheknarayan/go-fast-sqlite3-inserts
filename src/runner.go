package src

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/db"
)

type InsertionMethod uint32

const (
	M_Naive InsertionMethod = iota
	M_Naive_Async
	M_Naive_Prepared
	M_Naive_Async_Prepared
	M_Naive_Pragma_Optimized
	M_Batched_Optimized_Prepared
	M_Batched_Optimized_Prepared_Async
	// add more methods here
)

func (i InsertionMethod) ToString() string {
	switch i {
	case M_Naive:
		return "M_Naive"

	case M_Naive_Async:
		return "M_Naive_Async"
	case M_Naive_Prepared:
		return "M_Naive_Prepared"
	case M_Naive_Async_Prepared:
		return "M_Naive_Async_Prepared"
	case M_Naive_Pragma_Optimized:
		return "M_Naive_Pragma_Optimized"
	case M_Batched_Optimized_Prepared:
		return "M_Batched_Optimized_Prepared"
	case M_Batched_Optimized_Prepared_Async:
		return "M_Batched_Optimized_Prepared_Async"
	default:
		return "NULL"
	}
}

func Runner(method InsertionMethod, numberOfRows uint64, sqliteDB *sql.DB) {
	// down migrations once more
	db.RunMigrationDownScripts(sqliteDB)

	// up migrations here
	db.RunMigrationUpScripts(sqliteDB)

	if eq, rows := ValidateTable(0, sqliteDB); !eq {
		log.Panicf("Table not created properly, expected %d found %d", numberOfRows, rows)
	}

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
	case M_Naive_Prepared:
		{
			start := time.Now()
			NaivePrepared(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	case M_Naive_Pragma_Optimized:
		{
			start := time.Now()
			NaivePragmaOptimized(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	case M_Batched_Optimized_Prepared:
		{
			start := time.Now()
			BatchedPragmaOptimizedPrepared(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	case M_Batched_Optimized_Prepared_Async:
		{
			start := time.Now()
			BatchedPragmaOptimizedPreparedAsync(uint64(numberOfRows), sqliteDB)
			elapsed := time.Since(start)
			fmt.Printf("Time in seconds for %s: %f seconds \n", method.ToString(), elapsed.Seconds())
		}
	}
	// validate data
	if eq, rows := ValidateTable(numberOfRows, sqliteDB); !eq {
		log.Panicf("Rows not inserted properly, expected %d found %d", numberOfRows, rows)
	}

	// down migrations here
	err := db.RunMigrationDownScripts(sqliteDB)

	if err != nil {
		log.Panicf("Error while downing migrations: %s", err)
	}
}
