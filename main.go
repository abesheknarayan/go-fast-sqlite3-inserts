package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/db"
	"github.com/abesheknarayan/go-fast-sqlite-inserts/src"
)

func main() {
	numberOfRows := uint64(1e6 + 7)
	fmt.Printf("Number of rows: %d\n", numberOfRows)
	sqliteDB, err := db.NewDB("file:data/users.db")

	if err != nil {
		log.Panicf("Error while creating db instance %s", err)
	}

	src.Runner(src.M_Naive, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Naive_Async, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Naive_Prepared, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Naive_Pragma_Optimized, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Batched_Optimized_Prepared, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Batched_Optimized_Prepared_Async, uint64(numberOfRows), sqliteDB)

	// src.BatchedCheck(sqliteDB)

	defer func() {
		// run down migrations
		db.RunMigrationDownScripts(sqliteDB)
		sqliteDB.Close()
	}()

}
