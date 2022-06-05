package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/db"
	"github.com/abesheknarayan/go-fast-sqlite-inserts/src"
)

func main() {
	numberOfRows := 1e6
	sqliteDB, err := db.NewDB("file:data/users.db")

	if err != nil {
		log.Fatalf("Error while creating db instance %s", err)
	}

	err = db.RunMigrationScripts(sqliteDB)

	if err != nil {
		fmt.Printf("failed running migrations %s \n", err.Error())
	}

	src.Runner(src.M_Naive, uint64(numberOfRows), sqliteDB)

	src.Runner(src.M_Naive_Async, uint64(numberOfRows), sqliteDB)

}
