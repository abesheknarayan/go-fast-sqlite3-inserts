# go-fast-sqlite3-inserts
Playing around Golang + Sqlite3 to insert 100M rows as fast as possible

#### Status
Could insert upto 1e8 rows on my machine (with multiple crashes due to RAM overload). Will try to add more optimizations and documentation soon.

Checkout [Benchmarks for 1M rows](https://github.com/abesheknarayan/go-fast-sqlite3-inserts/blob/main/bench.txt) 

#### Resources
- [Avinassh Blog](https://avi.im/blag/2021/fast-sqlite-inserts/)
- [StackOverflow Question 1](https://stackoverflow.com/questions/1711631/improve-insert-per-second-performance-of-sqlite)
- [StackOverflow Question 2](https://stackoverflow.com/questions/12486436/how-do-i-batch-sql-statements-with-package-database-sql)
- [Sqlite Forum](https://sqlite.org/forum/info/f832398c19d30a4a)