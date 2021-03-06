package src

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/abesheknarayan/go-fast-sqlite-inserts/models"
)

// batched + pragma optimized + prepared without go-routines
func BatchedPragmaOptimizedPrepared(numberOfRows uint64, sqliteDB *sql.DB) {
	_, err := sqliteDB.Exec(`PRAGMA journal_mode = OFF;
    						 PRAGMA synchronous = 0; 
              				 PRAGMA cache_size = 1000000;
              				 PRAGMA locking_mode = EXCLUSIVE;
              				 PRAGMA temp_store = MEMORY;`)
	if err != nil {
		fmt.Printf("Error in setting pragmas %s", err)
	}

	// split the N rows into batches containing X number of rows --> N = X*k + R
	// have a single out channel where all these batch inserts preapred statements go
	// read from out channel and execute

	batchSize := 50
	numberOfBatches := int(math.Floor(float64(numberOfRows) / float64(batchSize)))

	queryString1 := make([]string, 0, batchSize)
	queryArgs1 := make([][]interface{}, 0, batchSize*4)

	for i := uint64(0); i < uint64(batchSize); i++ {
		queryString1 = append(queryString1, "(?,?,?,?)")
	}

	stmnt1 := fmt.Sprintf("insert into user(id,area,age,active) values %s", strings.Join(queryString1, ","))

	// except the last batch which may not be fully filled
	for i := uint64(0); i < uint64(numberOfBatches); i++ {
		// ith batch has contents from [ i*batchSize , (i + 1) * uint64(batchSize) - 1]
		l := i * uint64(batchSize)
		r := (i + 1) * uint64(batchSize)
		result := ComputeBatchAndReturnArgs(l, r)
		queryArgs1 = append(queryArgs1, result)
	}

	// last batch which contains remaining elements

	remaining := numberOfRows % uint64(batchSize)

	queryString2 := make([]string, 0, remaining)
	queryArgs2 := make([]interface{}, 0, remaining*4)

	for i := uint64(0); i < uint64(remaining); i++ {
		queryString2 = append(queryString2, "(?,?,?,?)")
		newUser := &models.User{
			Id:     uint32(i + 1 + uint64(numberOfBatches)*uint64(batchSize)),
			Area:   GenerateRandomAreaCode(),
			Age:    GenerateRandomAge(),
			Active: GenerateRandomBooleanInt(),
		}
		queryArgs2 = append(queryArgs2, newUser.Id)
		queryArgs2 = append(queryArgs2, newUser.Area)
		queryArgs2 = append(queryArgs2, newUser.Age)
		queryArgs2 = append(queryArgs2, newUser.Active)
	}

	stmnt2 := fmt.Sprintf("insert into user(id,area,age,active) values %s", strings.Join(queryString2, ","))

	// prepare everything and start tx
	tx, err := sqliteDB.Begin()

	if err != nil {
		log.Panicf(err.Error())
	}

	// prepare statement 1
	preparedStatement1, err := tx.Prepare(stmnt1)

	if err != nil {
		fmt.Println(err.Error())
	}

	for args := range queryArgs1 {
		_, err := preparedStatement1.Exec(queryArgs1[args]...)
		if err != nil {
			fmt.Println(err)
		}
	}

	// remaining

	if remaining > 0 {

		preparedStatement2, err := tx.Prepare(stmnt2)

		if err != nil {
			log.Panicf(err.Error())
		}
		_, err = preparedStatement2.Exec(queryArgs2...)

		if err != nil {
			log.Panicf(err.Error())
		}
	}

	tx.Commit()
}

func ComputeBatchAndReturnArgs(l uint64, r uint64) []interface{} {
	len := r - l
	queryArgs := make([]interface{}, 0, len*4)
	for i := l; i < r; i++ {
		newUser := &models.User{
			Id:     uint32(i + 1),
			Area:   GenerateRandomAreaCode(),
			Age:    GenerateRandomAge(),
			Active: GenerateRandomBooleanInt(),
		}
		// queryArgs = append(queryArgs, []interface{}{newUser.Id, newUser.Area, newUser.Age, newUser.Active})
		queryArgs = append(queryArgs, newUser.Id)
		queryArgs = append(queryArgs, newUser.Area)
		queryArgs = append(queryArgs, newUser.Age)
		queryArgs = append(queryArgs, newUser.Active)
	}
	return queryArgs
}
