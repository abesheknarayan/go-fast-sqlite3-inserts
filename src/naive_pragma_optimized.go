package src

import (
	"database/sql"
	"fmt"
)

func NaivePragmaOptimized(numberOfRows uint64, sqliteDB *sql.DB) {
	_, err := sqliteDB.Exec(`PRAGMA journal_mode = OFF;
    						 PRAGMA synchronous = 0; 
              				 PRAGMA cache_size = 1000000;
              				 PRAGMA locking_mode = EXCLUSIVE;
              				 PRAGMA temp_store = MEMORY;`)
	if err != nil {
		fmt.Printf("Error in setting pragmas %s", err)
	}
	NaivePrepared(numberOfRows, sqliteDB)
}
