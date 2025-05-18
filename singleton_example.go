package main

import (
	"fmt"
	"sync"
)

// Database represents a database connection
type Database struct {
	connection string
}

var (
	instance *Database
	once     sync.Once
)

// GetDatabase returns the single instance of Database
func GetDatabase() *Database {
	once.Do(func() {
		fmt.Println("Creating new database connection (this happens only once)")
		instance = &Database{
			connection: "mysql://localhost:3306",
		}
	})
	return instance
}

func (db *Database) Query(sql string) {
	fmt.Printf("Executing query '%s' on connection %s\n", sql, db.connection)
}

func main() {
	// Simulate multiple parts of the application trying to get the database
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Each goroutine tries to get the database instance
			db := GetDatabase()
			
			// All queries use the same connection
			db.Query(fmt.Sprintf("SELECT * FROM table_%d", id))
		}(i)
	}
	
	wg.Wait()
	
	// Prove it's the same instance
	db1 := GetDatabase()
	db2 := GetDatabase()
	fmt.Printf("\nAre db1 and db2 the same instance? %v\n", db1 == db2)
}
